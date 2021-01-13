package container

type Container interface {
	InitApp(filename string) error
	BuildUseCase(code string) (interface{}, error)
	Get(code string) (interface{}, bool)
	Put(code string, value interface{})
}
