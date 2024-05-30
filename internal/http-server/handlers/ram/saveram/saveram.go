package saveram

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

type RequestRAM struct {
	Name        string `json:"name" validate:"required"`
	Memory_type string `json:"memory_type" validate:"required"`
	Capacity    int64  `json:"capacity" validate:"required"`
}

type RAMSaver interface {
	SaveRAM(name, memory_type string, capacity int64) (int64, error)
}

func NewRAM(log *slog.Logger, ramSaver RAMSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.NewRAM"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req RequestRAM

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

		id, err := ramSaver.SaveRAM(req.Name, req.Memory_type, req.Capacity)
		if errors.Is(err, storage.ErrRAMAlreadyExists) {
			log.Info("ram already exists", slog.Int64("id", id))

			render.JSON(w, r, resp.Error("ram already exists"))

			return
		}

		if err != nil {
			log.Error("failed to save ram", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to save ram"))

			return
		}

		log.Info("ram saved", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
