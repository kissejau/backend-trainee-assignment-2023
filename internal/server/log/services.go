package log

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"time"
)

const (
	fileNamePrefix = "log"
	logsFolderName = "logs"
	URI            = "http://localhost:8080" // todo: hardcoded ;)
)

func GetLogFile(name string) (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path.Join(wd, logsFolderName, fmt.Sprintf("%v%v.csv", fileNamePrefix, name)))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GenerateLogLink(logs []Log) (*LogLinkDTO, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	time := time.Now().Format("2006010215040512")

	os.Mkdir(path.Join(wd, logsFolderName), 0755) // missing error (IF EXISTS)

	path := path.Join(wd, logsFolderName, fmt.Sprintf("%v%v.csv", fileNamePrefix, time))

	err = GenerateLog(path, logs)
	if err != nil {
		return nil, err
	}
	return &LogLinkDTO{fmt.Sprintf("%v%v/?log=%v", URI, LogUrl, time)}, nil
}

func GenerateLog(path string, logs []Log) error {
	csvData := [][]string{}
	for _, log := range logs {
		data := log.ToArray()
		csvData = append(csvData, data)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, data := range csvData {
		err := writer.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
