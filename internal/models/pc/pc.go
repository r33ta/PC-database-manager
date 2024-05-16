package pc

type PC struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	MemoryID  int64  `json:"memory_id"`
	CPUID     int64  `json:"cpu_id"`
	GPUID     int64  `json:"gpu_id"`
	StorageID int64  `json:"storage_id"`
}
