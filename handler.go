package zorro

type Handler interface {
	Handle(dto InputDto) (OutputDto, error)
	JobName() string
}
