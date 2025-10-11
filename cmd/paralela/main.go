package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/bernardothives/parallelCryptography/internal/crypto"
)

func main() {
	// 1. Setup - Validação de argumentos
	if len(os.Args) < 3 {
		fmt.Println("Uso: paralela <diretório_entrada> <diretório_saida> [num_goroutines]")
		fmt.Println("Exemplo: paralela ./assets/entrada ./assets/saida 4")
		fmt.Printf("Se num_goroutines não for especificado, usa %d (número de CPUs)\n", runtime.NumCPU())
		os.Exit(1)
	}

	dirEntrada := os.Args[1]
	dirSaida := os.Args[2]

	// Número de goroutines (padrão: número de CPUs)
	numGoroutines := runtime.NumCPU()
	if len(os.Args) >= 4 {
		var err error
		numGoroutines, err = strconv.Atoi(os.Args[3])
		if err != nil || numGoroutines < 1 {
			log.Fatalf("Número de goroutines inválido: %s", os.Args[3])
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
	fmt.Println("   PROCESSAMENTO PARALELO")
	fmt.Println("==========================================")
	fmt.Printf("Diretório de entrada: %s\n", dirEntrada)
	fmt.Printf("Diretório de saída: %s\n", dirSaida)
	fmt.Printf("Total de arquivos: %d\n", len(listaDeArquivos))
	fmt.Printf("Número de goroutines: %d\n", numGoroutines)
	fmt.Printf("CPUs disponíveis: %d\n", runtime.NumCPU())
	fmt.Println("------------------------------------------")
	fmt.Println("Iniciando processamento paralelo...")
	fmt.Println()

	// 3. Inicializar mecanismos de sincronização
	var wg sync.WaitGroup

	// Canal com buffer para controlar o número de goroutines simultâneas
	semaphore := make(chan struct{}, numGoroutines)

	// 4. Execução e Medição de Tempo
	tempoInicio := time.Now()

	// Dispara goroutines para processar arquivos em paralelo
	for i, nomeArquivo := range listaDeArquivos {
		wg.Add(1)

		// Adquire um slot no semáforo (limita goroutines simultâneas)
		semaphore <- struct{}{}

		// Inicia uma goroutine para processar o arquivo
		go func(index int, arquivo string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Libera o slot ao terminar

			caminhoEntrada := filepath.Join(dirEntrada, arquivo)
			caminhoSaida := filepath.Join(dirSaida, arquivo+".enc")

			fmt.Printf("[%d/%d] Processando: %s (goroutine)\n", index+1, len(listaDeArquivos), arquivo)

			// Criptografa o arquivo
			if err := crypto.CriptografarArquivo(caminhoEntrada, caminhoSaida, chaveCriptografia); err != nil {
				log.Printf("Erro ao criptografar arquivo %s: %v", arquivo, err)
			}
		}(i, nomeArquivo)
	}

	// Aguarda todas as goroutines terminarem
	wg.Wait()

	tempoFim := time.Now()
	duracaoTotal := tempoFim.Sub(tempoInicio)

	// 5. Resultado
	fmt.Println()
	fmt.Println("==========================================")
	fmt.Println("   RESULTADO")
	fmt.Println("==========================================")
	fmt.Println("Processamento paralelo concluído!")
	fmt.Printf("Tempo total de execução: %v\n", duracaoTotal)
	fmt.Printf("Tempo médio por arquivo: %v\n", duracaoTotal/time.Duration(len(listaDeArquivos)))
	fmt.Printf("Arquivos processados: %d\n", len(listaDeArquivos))
	fmt.Printf("Goroutines utilizadas: %d\n", numGoroutines)
	fmt.Println("==========================================")
}
