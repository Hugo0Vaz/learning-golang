package src

import (
	"fmt"
)

func List(path string) {
	queries, err := GetQueriesFromPath(path)

	if err != nil {
		fmt.Println("Erro buscando as querys: ", err)
	}

	for i, query := range queries {
		fmt.Println("Query num.: ", i)
		fmt.Println("Query: ", query.QueryString)
	}
}
