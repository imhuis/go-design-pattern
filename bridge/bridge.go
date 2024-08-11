package _bridge

type Interface interface {
	Operation() error
}

type Implementation struct {
	_ string
}

func NewImplementation() *Implementation {
	return &Implementation{}
}

func (i *Implementation) OperationImpl() error {
	// logic
	return nil
}

type Service interface {
	Do() error
}

type ServiceImpl struct {
	i Interface
}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) Do() error {
	return s.i.Operation()
}
