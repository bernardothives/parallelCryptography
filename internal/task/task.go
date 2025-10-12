package task

// Task representa uma unidade de trabalho a ser executada
type Task interface {
	Execute() error
	GetName() string
}

// SimpleTask é uma implementação básica de Task
type SimpleTask struct {
	Name string
	Fn   func() error
}

// Execute executa a tarefa
func (t *SimpleTask) Execute() error {
	return t.Fn()
}

// GetName retorna o nome da tarefa
func (t *SimpleTask) GetName() string {
	return t.Name
}

// NewTask cria uma nova tarefa simples
func NewTask(name string, fn func() error) *SimpleTask {
	return &SimpleTask{
		Name: name,
		Fn:   fn,
	}
}
