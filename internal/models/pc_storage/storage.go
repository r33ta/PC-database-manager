package pc_storage

const (
	SSD = "SSD"
	HDD = "HDD"
)

type Storage struct {
	ID         int
	Name       string
	Capacity   int
	MemoryType string
}
