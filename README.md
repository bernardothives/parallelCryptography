# Parallel Cryptography 🔐⚡

Projeto educacional para estudo de **computação paralela** em Go, implementando três versões de um sistema de criptografia de arquivos: sequencial, paralela e com biblioteca de pool de threads.

## 📚 Objetivo

Este projeto demonstra os conceitos de paralelismo através da implementação de diferentes abordagens para criptografar múltiplos arquivos, permitindo comparar:
- Performance sequencial vs paralela
- Ganhos de desempenho (speedup)
- Overhead de sincronização
- Padrões de design para computação paralela

## 🏗️ Estrutura do Projeto

```
parallelCryptography/
├── cmd/
│   ├── sequencial/     # Versão 1: Processamento sequencial
│   ├── paralela/       # Versão 2: Paralelismo com goroutines
│   └── biblioteca/     # Versão 3: Pool de workers (biblioteca)
├── internal/
│   ├── crypto/         # Funções de criptografia AES-GCM
│   └── task/           # Definições de tarefas
├── pkg/
│   └── parallel/       # Biblioteca de execução paralela
├── assets/
│   ├── entrada/        # Arquivos para criptografar
│   └── saida/          # Arquivos criptografados
└── compare.go          # Script de análise comparativa
```

## 🚀 Versões Implementadas

### 1️⃣ Versão Sequencial (`cmd/sequencial`)
**Algoritmo:** Processa arquivos um por um, de forma bloqueante.

**Características:**
- Implementação mais simples
- Serve como baseline para comparações
- Não utiliza paralelismo

**Como executar:**
```powershell
go run ./cmd/sequencial ./assets/entrada ./assets/saida
```

---

### 2️⃣ Versão Paralela (`cmd/paralela`)
**Algoritmo:** Usa goroutines e WaitGroup para paralelismo explícito.

**Características:**
- Extrai paralelismo através de múltiplas goroutines
- Usa semáforo para limitar goroutines simultâneas
- Controle manual de sincronização com `sync.WaitGroup`
- Permite especificar número de threads

**Como executar:**
```powershell
# Usa número de CPUs por padrão
go run ./cmd/paralela ./assets/entrada ./assets/saida

# Especifica 8 goroutines
go run ./cmd/paralela ./assets/entrada ./assets/saida 8
```

---

### 3️⃣ Versão com Biblioteca (`cmd/biblioteca`)
**Algoritmo:** Utiliza pool de workers (padrão Producer-Consumer).

**Características:**
- Biblioteca reutilizável de execução paralela
- Pool de workers pré-inicializado
- Padrão Producer-Consumer com canal bufferizado
- Interface simples: `Execute(task)`
- Abstração do paralelismo

**Como executar:**
```powershell
# Usa número de CPUs por padrão
go run ./cmd/biblioteca ./assets/entrada ./assets/saida

# Especifica 4 workers
go run ./cmd/biblioteca ./assets/entrada ./assets/saida 4
```

---

## 📊 Análise Comparativa

Execute todas as versões e compare os resultados:

```powershell
go run compare.go ./assets/entrada ./assets/saida
```

### Saída Exemplo:
```
╔════════════════════════════════════════════════════════╗
║  ANÁLISE COMPARATIVA DE PERFORMANCE                    ║
║  Criptografia de Arquivos - Paralela vs Sequencial    ║
╚════════════════════════════════════════════════════════╝

📁 Total de arquivos: 10
💻 CPUs disponíveis: 8

┌────────────────────────────┬──────────┬──────────────┬──────────────┐
│ Versão                     │ Workers  │ Tempo Total  │ Tempo/Arq    │
├────────────────────────────┼──────────┼──────────────┼──────────────┤
│ Sequencial                 │        1 │      2.456s  │      245ms   │
│ Paralela (4 threads)       │        4 │      782ms   │       78ms   │
│ Biblioteca (8 workers)     │        8 │      524ms   │       52ms   │
└────────────────────────────┴──────────┴──────────────┴──────────────┘

📊 SPEEDUP:
   Sequencial: 1.00x (baseline)
   Paralela (4 threads): 3.14x (68.2% mais rápido)
   Biblioteca (8 workers): 4.69x (78.7% mais rápido)
```

## 🔧 Biblioteca de Execução Paralela

A biblioteca (`pkg/parallel/executor.go`) oferece uma API simples:

```go
// Criar executor com 4 workers e buffer de 100 tarefas
executor := parallel.NewExecutor(4, 100)

// Submeter tarefas
executor.Execute(func() {
    // Sua tarefa aqui
})

// Aguardar conclusão
executor.Wait()
```

**Padrões implementados:**
- ✅ Producer-Consumer Pattern
- ✅ Worker Pool Pattern
- ✅ Channel-based Communication
- ✅ Graceful Shutdown

## 🧪 Tecnologias Utilizadas

- **Linguagem:** Go 1.25.2
- **Criptografia:** AES-256-GCM
- **Paralelismo:** Goroutines, Channels, sync.WaitGroup
- **Padrões:** Pool de Threads, Producer-Consumer

## 📈 Vantagens do Paralelismo

✅ **Performance:**
- Redução significativa do tempo de processamento
- Melhor aproveitamento de CPUs multi-core
- Speedup próximo ao linear em operações CPU-bound

✅ **Escalabilidade:**
- Processa grandes volumes de arquivos eficientemente
- Adapta-se ao hardware disponível

✅ **Modularidade:**
- Biblioteca reutilizável
- Separação de responsabilidades

## ⚠️ Desvantagens e Considerações

❌ **Complexidade:**
- Código mais complexo
- Necessidade de sincronização
- Possíveis race conditions

❌ **Overhead:**
- Criação e gerenciamento de goroutines
- Contenção em I/O (disco)
- Não há ganho se I/O for o gargalo

❌ **Debugging:**
- Mais difícil de debugar
- Comportamento não-determinístico

## 🎯 Conceitos Aprendidos

1. **Paralelismo vs Concorrência**
   - Diferença entre execução simultânea e gestão de múltiplas tarefas

2. **Goroutines e Channels**
   - Modelo de concorrência do Go (CSP)
   - Comunicação através de canais

3. **Sincronização**
   - WaitGroups para aguardar conclusão
   - Semáforos para limitar recursos

4. **Padrões de Design**
   - Worker Pool
   - Producer-Consumer
   - Task Queue

5. **Medição de Performance**
   - Benchmarking
   - Cálculo de speedup
   - Análise de gargalos

## 🚀 Como Começar

1. **Clone o repositório:**
```powershell
git clone https://github.com/bernardothives/parallelCryptography.git
cd parallelCryptography
```

2. **Crie alguns arquivos de teste:**
```powershell
# Cria diretórios
New-Item -ItemType Directory -Force -Path assets/entrada
New-Item -ItemType Directory -Force -Path assets/saida

# Cria arquivos de teste (Windows PowerShell)
1..10 | ForEach-Object { 
    "Conteúdo do arquivo $_" * 1000 | Out-File "assets/entrada/arquivo$_.txt" 
}
```

3. **Execute as versões:**
```powershell
# Sequencial
go run ./cmd/sequencial ./assets/entrada ./assets/saida

# Paralela
go run ./cmd/paralela ./assets/entrada ./assets/saida 4

# Biblioteca
go run ./cmd/biblioteca ./assets/entrada ./assets/saida 4
```

4. **Compare os resultados:**
```powershell
go run compare.go ./assets/entrada ./assets/saida
```

## 📝 Notas de Segurança

⚠️ **ATENÇÃO:** Este projeto é para fins educacionais. A chave de criptografia está hardcoded no código para simplicidade. Em produção:
- Use chaves geradas de forma segura
- Armazene chaves em local seguro (vault, HSM)
- Use KDF (Key Derivation Function) adequada
- Implemente rotação de chaves

## 📄 Licença

Este projeto é de código aberto e está disponível para fins educacionais.

## 👨‍💻 Autor

Desenvolvido para estudo de computação paralela em Go.

---

**Happy Parallel Computing! 🚀⚡**