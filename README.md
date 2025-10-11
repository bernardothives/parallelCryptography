## Estrutura do Projeto

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

### Versão Sequencial (`cmd/sequencial`)
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

### Versão Paralela (`cmd/paralela`)
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

### Versão com Biblioteca (`cmd/biblioteca`)
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


## Padrões implementados:
- ✅ Producer-Consumer Pattern
- ✅ Worker Pool Pattern
- ✅ Channel-based Communication
- ✅ Graceful Shutdown
