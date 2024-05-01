package src

import (
	"fmt"
)

func RunTests(host string,
	port string,
	user string,
	pass string,
	path string,
	database string,
	dbtype string,
	reps int) {

	// queries, err := GetQueriesFromPath(path)
	//
	// if err != nil {
	// 	fmt.Println("ERROR GETTING THE QUERIES: ", err)
	// }

	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("User:", user)
	fmt.Println("Pass:", pass)
	fmt.Println("Path:", path)
	fmt.Println("Database:", database)
	fmt.Println("DBType:", dbtype)
}
