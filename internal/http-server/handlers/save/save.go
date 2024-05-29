package save

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

type RequestRAM struct {
	Name        string `json:"name" validate:"required"`
	Memory_type string `json:"memory_type" validate:"required"`
	Capacity    int64  `json:"capacity" validate:"required"`
}

type RequestCPU struct {
	Name      string `json:"name" validate:"required"`
	Cores     int64  `json:"cores" validate:"required"`
	Threads   int64  `json:"threads" validate:"required"`
	Frequency int64  `json:"frequency" validate:"required"`
}

type RequestGPU struct {
	Name         string `json:"name" validate:"required"`
	Manufacturer string `json:"manufacturer" validate:"required"`
	Memory       int64  `json:"memory" validate:"required"`
	Frequency    int64  `json:"frequency" validate:"required"`
}

type RequestMemory struct {
	Name        string `json:"name" validate:"required"`
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
