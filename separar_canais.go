package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Uso: go run script.go <arquivo_log> <canal_inicial> <canal_final>")
		os.Exit(1)
	}

	arquivoLog := os.Args[1]
	canalInicial := os.Args[2]
	canalFinal := os.Args[3]

	pastaDestino := "logs_canais"
	if _, err := os.Stat(pastaDestino); os.IsNotExist(err) {
		os.Mkdir(pastaDestino, 0755)
	}

	for i := canalInicialInt(canalInicial); i <= canalInicialInt(canalFinal); i++ {
		canal := fmt.Sprintf("C%02d", i)
		nomeArquivo := fmt.Sprintf("%s_%s.txt", "logs_canal", canal)
		caminhoArquivo := filepath.Join(pastaDestino, nomeArquivo)

		cmd := exec.Command("grep", "-R", canal, arquivoLog)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("Erro ao executar o comando grep:", err)
			os.Exit(1)
		}

		f, err := os.OpenFile(caminhoArquivo, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Erro ao criar ou abrir o arquivo:", err)
			os.Exit(1)
		}

		defer f.Close()
		f.WriteString(string(out))
	}

	fmt.Println("Processo concluído. Os logs foram salvos na pasta", pastaDestino)
}

func canalInicialInt(canalInicial string) int {
	var canalInt int
	fmt.Sscanf(canalInicial, "C%d", &canalInt)
	return canalInt
}
