package memory

const (
	SSD = "SSD"
	HDD = "HDD"
)

type Memory struct {
	ID          int64
	Name        string
	Capacity    int64
	StorageType string
}
