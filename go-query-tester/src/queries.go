package src

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
)

type Query struct {
	Comment     string
	QueryString string
}

func GetQueriesFromPath(path string) ([]Query, error) {
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

	var queries []Query

	for scanner.Scan() {
		line := scanner.Text()

		if !getLineMatch(line) {
			continue
		}

		queries = append(queries, Query{Comment: "...", QueryString: line})

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return queries, nil
}

func getLineMatch(line string) bool {
	queryPattern := `(?i)^\s*(SELECT|INSERT|UPDATE|DELETE|CREATE|ALTER|DROP|TRUNCATE|REPLACE|MERGE).+?;`
	queryRegexp := regexp.MustCompile(queryPattern)

	return queryRegexp.MatchString(line)
}
