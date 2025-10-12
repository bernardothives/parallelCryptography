package parallel

import (
	"sync"
)

// Task representa uma tarefa a ser executada (função sem parâmetros)
type Task func()

// Executor gerencia um pool de workers para execução paralela de tarefas
type Executor struct {
	taskChannel chan Task
	wg          sync.WaitGroup
	numWorkers  int
}

// NewExecutor cria um novo executor com um pool de workers
// numWorkers: número de goroutines workers no pool
// bufferSize: tamanho do buffer do canal de tarefas (0 para canal não-bufferizado)
func NewExecutor(numWorkers int, bufferSize int) *Executor {
	executor := &Executor{
		taskChannel: make(chan Task, bufferSize),
		numWorkers:  numWorkers,
	}

	// Inicia os workers (consumidores)
	for i := 0; i < numWorkers; i++ {
		executor.wg.Add(1)
		go executor.worker()
	}

	return executor
}

// worker é a goroutine que consome e executa tarefas do canal
func (e *Executor) worker() {
	defer e.wg.Done()
	
	// Consome tarefas do canal até que ele seja fechado
	for task := range e.taskChannel {
		task() // Executa a tarefa
	}
}

// Execute submete uma nova tarefa para ser executada pelo pool
// Esta função não bloqueia se o canal tiver buffer disponível
func (e *Executor) Execute(task Task) {
	e.taskChannel <- task
}

// Wait aguarda a conclusão de todas as tarefas
// Fecha o canal de tarefas e aguarda todos os workers terminarem
func (e *Executor) Wait() {
	// Fecha o canal, sinalizando que não haverá mais tarefas
	close(e.taskChannel)
	
	// Aguarda todos os workers terminarem
	e.wg.Wait()
}

// GetWorkerCount retorna o número de workers no pool
func (e *Executor) GetWorkerCount() int {
	return e.numWorkers
}
