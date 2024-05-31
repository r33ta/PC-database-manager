package savepc

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	resp "github.com/r33ta/pc-database-manager/internal/lib/api/response"
	"github.com/r33ta/pc-database-manager/internal/lib/logger/sl"
	"github.com/r33ta/pc-database-manager/internal/storage"
)

type Response struct {
	resp.Response
}

type RequestPC struct {
	Name     string `json:"name" validate:"required"`
	RAMID    int64  `json:"ram_id" validate:"required"`
	CPUID    int64  `json:"cpu_id" validate:"required"`
	GPUID    int64  `json:"gpu_id" validate:"required"`
	MemoryID int64  `json:"memory_id" validate:"required"`
}

type PCSaver interface {
	SavePC(name string, ramID, cpuID, gpuID, memoryID int64) (int64, error)
}

func New(log *slog.Logger, pcSaver PCSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.savepc.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req RequestPC

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request body"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validateErr))

			return
		}

		id, err := pcSaver.SavePC(req.Name, req.RAMID, req.CPUID, req.GPUID, req.MemoryID)
		if errors.Is(err, storage.ErrPCAlreadyExists) {
			log.Info("pc already exists", slog.Int64("id", id))

			render.JSON(w, r, resp.Error("pc already exists"))

			return
		}

		if err != nil {
			log.Error("failed to save pc", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to save pc"))

			return
		}

		log.Info("pc saved", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
