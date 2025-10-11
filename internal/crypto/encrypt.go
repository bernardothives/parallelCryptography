package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func CriptografarArquivo(caminhoEntrada, caminhoSaida string, chave []byte) erro {
	bytesArquivoOriginal, err := os.ReadFile(caminhoEntrada)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(chave)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	dadosCriptografados := gcm.Seal(nonce, nonce, bytesArquivoOriginal, nil)

	err = os.WriteFile(caminhoSaida, dadosCriptografados, 0644)
	if err != nil {
		return err
	}
	return nil
}
