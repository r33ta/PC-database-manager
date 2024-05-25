package save

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	resp "github.com/r33ta/pc-database-manager/internal/lib/api/response"
	"github.com/r33ta/pc-database-manager/internal/lib/logger/sl"
)

type Request struct {
	Name     string `json:"name"`
	RAMID    int64  `json:"ram_id" validate:"required"`
	CPUID    int64  `json:"cpu_id" validate:"required"`
	GPUID    int64  `json:"gpu_id" validate:"required"`
	MemoryID int64  `json:"memory_id" validate:"required"`
}

type Response struct {
	resp.Response
	Name string `json:"name"`
}

type PCSaver interface {
	SavePC(name string, ramID, cpuID, gpuID, memoryID int64) (int64, error)
}

func New(log *slog.Logger, pcSaver PCSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request body"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.Error("invalid request"))

			return
		}
	}
}
