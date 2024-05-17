package ram

const (
	DDR3 = "DDR3"
	DDR4 = "DDR4"
	DDR5 = "DDR5"
)

type RAM struct {
	ID         int64
	Name       string
	MemoryType string
	Capacity   int64
}
