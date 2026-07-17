package zorro

type Handler interface {
	Handle(dto InputDto) ([]Variable, error)
	JobName() string
}
