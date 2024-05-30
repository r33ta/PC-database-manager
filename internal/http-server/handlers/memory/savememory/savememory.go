package savememory

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

type RequestMemory struct {
	Name        string `json:"name" validate:"required"`
	Capacity    int64  `json:"capacity" validate:"required"`
	StorageType string `json:"storage_type" validate:"required"`
}

type MemorySaver interface {
	SaveMemory(name string, capacity int64, storage_type string) (int64, error)
}

func NewMemory(log *slog.Logger, memorySaver MemorySaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.NewMemory"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req RequestMemory

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

		id, err := memorySaver.SaveMemory(req.Name, req.Capacity, req.StorageType)
		if errors.Is(err, storage.ErrMemoryAlreadyExists) {
			log.Info("memory already exists", slog.Int64("id", id))

			render.JSON(w, r, resp.Error("memory already exists"))

			return
		}

		if err != nil {
			log.Error("failed to save memory", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to save memory"))

			return
		}

		log.Info("memory saved", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
