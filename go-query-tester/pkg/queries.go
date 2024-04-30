package pkg

import (
	"bufio"
	"os"
	"path/filepath"
)

type Query struct {
	Comment     string
	QueryString string
}

func getQueriesFromPath(path string) ([]Query, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		queries, err := getQueriesFromDir(path)

		if err != nil {
			return nil, err
		}

		return queries, nil

	} else {
		queries, err := getQueriesFromFile(path)

		if err != nil {
			return nil, err
		}

		return queries, nil
	}
}

func getQueriesFromDir(path string) ([]Query, error) {
	var queries []Query

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		fileExtension := filepath.Ext(e.Name())

		if fileExtension != ".sql" {
			continue
		}

		fileQueries, err := getQueriesFromFile(e.Name())

		if err != nil {
			queries = append(queries, fileQueries...)
		}
	}

	return queries, nil
}

func getQueriesFromFile(path string) ([]Query, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// not reconized linetype -> 0
	// comment line linetype -> 1
	// query line linetype -> 2
	lastLineType := 0
	lastLineContent := ""

	for scanner.Scan() {
		line := scanner.Text()

		lineType, lineContent, err := parseLine(line)

		if err != nil {
			continue
		}

		if lineType == 1 && lastLineType == 1 {
			lastLineContent := append(lastLineContent, lineContent)
		} else if lineType == 2 && lastLineType == 2 {
			lastLineContent := append(lastLineContent, lineContent)
		} else if lineType == 2 && lastLineType == 1 {

		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}
