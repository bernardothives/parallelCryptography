package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/bernardothives/parallelCryptography/internal/crypto"
	"github.com/bernardothives/parallelCryptography/pkg/parallel"
)

func main() {
	// 1. Setup - Validação de argumentos
	if len(os.Args) < 3 {
		fmt.Println("Uso: biblioteca <diretório_entrada> <diretório_saida> [num_workers]")
		fmt.Println("Exemplo: biblioteca ./assets/entrada ./assets/saida 4")
		fmt.Printf("Se num_workers não for especificado, usa %d (número de CPUs)\n", runtime.NumCPU())
		os.Exit(1)
	}

	dirEntrada := os.Args[1]
	dirSaida := os.Args[2]

	// Número de workers no pool (padrão: número de CPUs)
	numWorkers := runtime.NumCPU()
	if len(os.Args) >= 4 {
		var err error
		numWorkers, err = strconv.Atoi(os.Args[3])
		if err != nil || numWorkers < 1 {
			log.Fatalf("Número de workers inválido: %s", os.Args[3])
		}
	}

	// Gera/carrega chave fixa de criptografia
	chaveCriptografia := crypto.GerarChaveFixa()

	// Cria o diretório de saída se não existir
	if err := os.MkdirAll(dirSaida, 0755); err != nil {
		log.Fatalf("Erro ao criar diretório de saída: %v", err)
	}

	// 2. Listar arquivos do diretório de entrada
	arquivos, err := os.ReadDir(dirEntrada)
	if err != nil {
		log.Fatalf("Erro ao listar diretório de entrada: %v", err)
	}

	// Filtra apenas arquivos regulares
	var listaDeArquivos []string
	for _, arquivo := range arquivos {
		if !arquivo.IsDir() {
			listaDeArquivos = append(listaDeArquivos, arquivo.Name())
		}
	}

	if len(listaDeArquivos) == 0 {
		fmt.Println("Nenhum arquivo encontrado no diretório de entrada.")
		return
	}

	fmt.Println("==========================================")
	fmt.Println("   PROCESSAMENTO COM BIBLIOTECA PARALELA")
	fmt.Println("==========================================")
	fmt.Printf("Diretório de entrada: %s\n", dirEntrada)
	fmt.Printf("Diretório de saída: %s\n", dirSaida)
	fmt.Printf("Total de arquivos: %d\n", len(listaDeArquivos))
	fmt.Printf("Workers no pool: %d\n", numWorkers)
	fmt.Printf("CPUs disponíveis: %d\n", runtime.NumCPU())
	fmt.Println("------------------------------------------")
	fmt.Println("Iniciando processamento com biblioteca paralela...")
	fmt.Println()

	// 3. Criar o Executor (pool de threads)
	// Buffer de 100 permite enfileirar tarefas sem bloqueio
	executor := parallel.NewExecutor(numWorkers, 100)

	// 4. Execução e Medição de Tempo
	tempoInicio := time.Now()

	// Submete todas as tarefas ao executor (produtor)
	for i, nomeArquivo := range listaDeArquivos {
		// Captura as variáveis para uso dentro da closure
		index := i
		arquivo := nomeArquivo

		// Cria uma tarefa (função) para criptografar o arquivo
		tarefa := func() {
			caminhoEntrada := filepath.Join(dirEntrada, arquivo)
			caminhoSaida := filepath.Join(dirSaida, arquivo+".enc")

			fmt.Printf("[%d/%d] Processando: %s (worker pool)\n", index+1, len(listaDeArquivos), arquivo)

			// Criptografa o arquivo
			if err := crypto.CriptografarArquivo(caminhoEntrada, caminhoSaida, chaveCriptografia); err != nil {
				log.Printf("Erro ao criptografar arquivo %s: %v", arquivo, err)
			}
		}

		// Submete a tarefa ao executor
		executor.Execute(tarefa)
	}

	// Aguarda todas as tarefas serem concluídas
	executor.Wait()

	tempoFim := time.Now()
	duracaoTotal := tempoFim.Sub(tempoInicio)

	// 5. Resultado
	fmt.Println()
	fmt.Println("==========================================")
	fmt.Println("   RESULTADO")
	fmt.Println("==========================================")
	fmt.Println("✓ Processamento com biblioteca concluído!")
	fmt.Printf("⏱  Tempo total de execução: %v\n", duracaoTotal)
	fmt.Printf("📊 Tempo médio por arquivo: %v\n", duracaoTotal/time.Duration(len(listaDeArquivos)))
	fmt.Printf("📁 Arquivos processados: %d\n", len(listaDeArquivos))
	fmt.Printf("🔧 Workers utilizados: %d\n", executor.GetWorkerCount())
	fmt.Println("==========================================")
}
