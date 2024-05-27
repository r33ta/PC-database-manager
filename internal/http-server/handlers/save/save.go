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

type Response struct {
	resp.Response
	Name string `json:"name"`
}

type RequestPC struct {
	Name     string `json:"name"`
	RAMID    int64  `json:"ram_id" validate:"required"`
	CPUID    int64  `json:"cpu_id" validate:"required"`
	GPUID    int64  `json:"gpu_id" validate:"required"`
	MemoryID int64  `json:"memory_id" validate:"required"`
}

type RequestRAM struct {
	Name        string `json:"name"`
	Memory_type string `json:"memory_type" validate:"required"`
	Capacity    int64  `json:"capacity" validate:"required"`
}

type RequestCPU struct {
	Name      string `json:"name"`
	Cores     int64  `json:"cores" validate:"required"`
	Threads   int64  `json:"threads" validate:"required"`
	Frequency int64  `json:"frequency" validate:"required"`
}

type RequestGPU struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer" validate:"required"`
	Memory       int64  `json:"memory" validate:"required"`
	Frequency    int64  `json:"frequency" validate:"required"`
}

type RequestMemory struct {
	Name        string `json:"name"`
	Capacity    int64  `json:"capacity" validate:"required"`
	StorageType string `json:"storage_type" validate:"required"`
}

type PCSaver interface {
	SavePC(name string, ramID, cpuID, gpuID, memoryID int64) (int64, error)
}

type RAMSaver interface {
	SaveRAM(name, memory_type string, capacity int64) (int64, error)
}

type CPUSaver interface {
	SaveCPU(name string, cores, threads, frequency int64) (int64, error)
}

type GPUSaver interface {
	SaveGPU(name, manufacturer string, memory, frequency int64) (int64, error)
}

type MemorySaver interface {
	SaveMemory(name string, capacity int64, storage_type string) (int64, error)
}

func NewPC(log *slog.Logger, pcSaver PCSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.NewPC"

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

		name := req.Name
		if name == "" {
			name = "Unknown PC"
		}
	}
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

		name := req.Name
		if name == "" {
			name = "Unknown RAM"
		}
	}
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

		name := req.Name
		if name == "" {
			name = "Unknown CPU"
		}
	}
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

		name := req.Name
		if name == "" {
			name = "Unknown GPU"
		}
	}
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

		name := req.Name
		if name == "" {
			name = "Unknown Memory"
		}
	}
}
