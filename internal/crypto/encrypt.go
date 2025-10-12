package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

// CriptografarArquivo criptografa um arquivo usando AES-GCM
// caminhoEntrada: caminho do arquivo a ser criptografado
// caminhoSaida: caminho onde o arquivo criptografado será salvo
// chave: chave de criptografia (deve ter 16, 24 ou 32 bytes)
func CriptografarArquivo(caminhoEntrada, caminhoSaida string, chave []byte) error {
	// Lê o conteúdo do arquivo original
	bytesArquivoOriginal, err := os.ReadFile(caminhoEntrada)
	if err != nil {
		return err
	}

	// Cria o cipher block AES
	block, err := aes.NewCipher(chave)
	if err != nil {
		return err
	}

	// Cria o GCM (Galois/Counter Mode)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Gera um nonce aleatório
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Criptografa os dados
	dadosCriptografados := gcm.Seal(nonce, nonce, bytesArquivoOriginal, nil)

	// Escreve o arquivo criptografado
	err = os.WriteFile(caminhoSaida, dadosCriptografados, 0644)
	if err != nil {
		return err
	}
	
	return nil
}

// GerarChaveFixa gera uma chave fixa de 32 bytes para AES-256
// Em produção, use uma chave segura e armazenada adequadamente
func GerarChaveFixa() []byte {
	// Chave fixa de exemplo (32 bytes para AES-256)
	// ATENÇÃO: Em produção, use uma chave gerada de forma segura!
	return []byte("chave-super-secreta-de-32-bytes!")
}

