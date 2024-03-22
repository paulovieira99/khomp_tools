package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Uso: go run script.go <diretorio_logs> <data_hora_inicial> <data_hora_final>")
		os.Exit(1)
	}

	logDirectory := os.Args[1]
	startTimeStr := os.Args[2]
	endTimeStr := os.Args[3]

	startTime, err := time.Parse("02/01/2006 15:04:05.999", startTimeStr)
	if err != nil {
		fmt.Println("Erro ao analisar a data/hora inicial:", err)
		os.Exit(1)
	}

	endTime, err := time.Parse("02/01/2006 15:04:05.999", endTimeStr)
	if err != nil {
		fmt.Println("Erro ao analisar a data/hora final:", err)
		os.Exit(1)
	}

	outputFileName := "logs_selecionados.txt"

	if err := extractLogs(logDirectory, startTime, endTime, outputFileName); err != nil {
		fmt.Println("Erro ao extrair logs:", err)
		os.Exit(1)
	}

	fmt.Printf("Logs no intervalo de %s a %s foram extraídos para %s\n", startTime, endTime, outputFileName)
}

func extractLogs(logDirectory string, startTime, endTime time.Time, outputFile string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	logFiles := []string{"k3l_intf.log", "kgateway.log", "voip_k3l.log", "r2.log", "isdn_k3l.log", "voip_kmp.log"}

	for _, logFile := range logFiles {
		logFilePath := logDirectory + "/" + logFile
		if fileExists(logFilePath) {
			_, _ = output.WriteString(fmt.Sprintf("==== %s ====\n\n", logFile))
			if err := extractLogsFromFile(logFilePath, startTime, endTime, output); err != nil {
				return err
			}
			_, _ = output.WriteString("\n\n")
		}
	}

	return nil
}

func extractLogsFromFile(logFilePath string, startTime, endTime time.Time, output *os.File) error {
	file, err := os.Open(logFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		timestampMatch := regexp.MustCompile(`\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2}.\d{3}`).FindString(line)
		if timestampMatch != "" {
			logTimestamp, err := time.Parse("02/01/2006 15:04:05.999", timestampMatch)
			if err != nil {
				fmt.Printf("Aviso: Não foi possível analisar a data/hora da linha: %s\n", line)
				continue
			}

			if logTimestamp.After(startTime) && logTimestamp.Before(endTime) {
				_, err := output.WriteString(line + "\n")
				if err != nil {
					return err
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
