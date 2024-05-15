package memory

const (
	DDR3 = "DDR3"
	DDR4 = "DDR4"
	DDR5 = "DDR5"
)

type Memory struct {
	ID         int
	Name       string
	MemoryType string
	Capacity   int
}
