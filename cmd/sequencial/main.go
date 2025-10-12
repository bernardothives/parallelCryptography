package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/bernardothives/parallelCryptography/internal/crypto"
)

func main() {
	// 1. Setup - Validação de argumentos
	if len(os.Args) < 3 {
		fmt.Println("Uso: sequencial <diretório_entrada> <diretório_saida>")
		fmt.Println("Exemplo: sequencial ./assets/entrada ./assets/saida")
		os.Exit(1)
	}

	dirEntrada := os.Args[1]
	dirSaida := os.Args[2]

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

	// Filtra apenas arquivos regulares (não diretórios)
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
	fmt.Println("   PROCESSAMENTO SEQUENCIAL")
	fmt.Println("==========================================")
	fmt.Printf("Diretório de entrada: %s\n", dirEntrada)
	fmt.Printf("Diretório de saída: %s\n", dirSaida)
	fmt.Printf("Total de arquivos: %d\n", len(listaDeArquivos))
	fmt.Println("------------------------------------------")
	fmt.Println("Iniciando processamento sequencial...")
	fmt.Println()

	// 3. Execução e Medição de Tempo
	tempoInicio := time.Now()

	// Processa cada arquivo sequencialmente (um por vez)
	for i, nomeArquivo := range listaDeArquivos {
		caminhoEntrada := filepath.Join(dirEntrada, nomeArquivo)
		caminhoSaida := filepath.Join(dirSaida, nomeArquivo+".enc")

		fmt.Printf("[%d/%d] Processando: %s\n", i+1, len(listaDeArquivos), nomeArquivo)

		// Criptografa o arquivo (bloqueante/sequencial)
		if err := crypto.CriptografarArquivo(caminhoEntrada, caminhoSaida, chaveCriptografia); err != nil {
			log.Printf("Erro ao criptografar arquivo %s: %v", nomeArquivo, err)
			continue
		}
	}

	tempoFim := time.Now()
	duracaoTotal := tempoFim.Sub(tempoInicio)

	// 4. Resultado
	fmt.Println()
	fmt.Println("==========================================")
	fmt.Println("   RESULTADO")
	fmt.Println("==========================================")
	fmt.Println("Processamento sequencial concluído!")
	fmt.Printf("Tempo total de execução: %v\n", duracaoTotal)
	fmt.Printf("Tempo médio por arquivo: %v\n", duracaoTotal/time.Duration(len(listaDeArquivos)))
	fmt.Printf("Arquivos processados: %d\n", len(listaDeArquivos))
	fmt.Println("==========================================")
}
