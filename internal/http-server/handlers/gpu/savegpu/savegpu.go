package savegpu

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

type RequestGPU struct {
	Name         string `json:"name" validate:"required"`
	Manufacturer string `json:"manufacturer" validate:"required"`
	Memory       int64  `json:"memory" validate:"required"`
	Frequency    int64  `json:"frequency" validate:"required"`
}

type GPUSaver interface {
	SaveGPU(name, manufacturer string, memory, frequency int64) (int64, error)
}

func NewGPU(log *slog.Logger, gpuSaver GPUSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.NewGPU"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req RequestGPU

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

		id, err := gpuSaver.SaveGPU(req.Name, req.Manufacturer, req.Memory, req.Frequency)
		if errors.Is(err, storage.ErrGPUAlreadyExists) {
			log.Info("gpu already exists", slog.Int64("id", id))

			render.JSON(w, r, resp.Error("gpu already exists"))

			return
		}

		if err != nil {
			log.Error("failed to save gpu", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to save gpu"))

			return
		}

		log.Info("gpu saved", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
