package savecpu

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

type RequestCPU struct {
	Name      string `json:"name" validate:"required"`
	Cores     int64  `json:"cores" validate:"required"`
	Threads   int64  `json:"threads" validate:"required"`
	Frequency int64  `json:"frequency" validate:"required"`
}

type CPUSaver interface {
	SaveCPU(name string, cores, threads, frequency int64) (int64, error)
}

func NewCPU(log *slog.Logger, cpuSaver CPUSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.NewCPU"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req RequestCPU

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

		id, err := cpuSaver.SaveCPU(req.Name, req.Cores, req.Threads, req.Frequency)
		if errors.Is(err, storage.ErrCPUAlreadyExists) {
			log.Info("cpu already exists", slog.Int64("id", id))

			render.JSON(w, r, resp.Error("cpu already exists"))

			return
		}

		if err != nil {
			log.Error("failed to save cpu", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to save cpu"))

			return
		}

		log.Info("cpu saved", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
