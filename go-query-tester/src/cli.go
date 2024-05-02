package src

import (
	"flag"
	"fmt"
	"os"
)

func ParseArgs() {
	testCmdFlagSet := flag.NewFlagSet("test", flag.ExitOnError)
	testHostFlag := testCmdFlagSet.String("host", "localhost", "the database host")
	testPortFlag := testCmdFlagSet.String("port", "42069", "the database port")
	testUserFlag := testCmdFlagSet.String("user", "root", "the database user")
	testPassFlag := testCmdFlagSet.String("pass", "password", "the database password")
	testPathFlag := testCmdFlagSet.String("path", "queries", "the queries file or directory")
	testDatabaseFlag := testCmdFlagSet.String("db", "table", "the database to test")
	testDBTypeFlag := testCmdFlagSet.String("type", "mysql", "the type of database to be tested")
	testRepetitionFlag := testCmdFlagSet.Int("nr", 1, "the number of repetitions of each query")

	listCmdFlagSet := flag.NewFlagSet("list", flag.ExitOnError)
	listPathFlag := listCmdFlagSet.String("path", "queries", "the queries file or directory")

	runCmdFlagSet := flag.NewFlagSet("run", flag.ExitOnError)
	runHostFlag := runCmdFlagSet.String("host", "localhost", "the database host")
	runPortFlag := runCmdFlagSet.String("port", "42069", "the database port")
	runUserFlag := runCmdFlagSet.String("user", "root", "the database user")
	runPassFlag := runCmdFlagSet.String("pass", "password", "the database password")
	runPathFlag := runCmdFlagSet.String("path", "queries", "the queries file or directory")
	runDatabaseFlag := runCmdFlagSet.String("db", "table", "the database to test")
	runDBTypeFlag := runCmdFlagSet.String("type", "mysql", "the type of database to be tested")
	runRepetitionFlag := runCmdFlagSet.Int("nr", 1, "the number of repetitions of each query")

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		fmt.Println("testing connections and queries...")
		fmt.Println("\n ")
		runCmdFlagSet.Parse(os.Args[2:])
		testConnection(*runHostFlag,
			*runPortFlag,
			*runUserFlag,
			*runPassFlag,
			*runPathFlag,
			*runDatabaseFlag,
			*runDBTypeFlag,
			*runRepetitionFlag)
	case "test":
		fmt.Println("running tests...")
		fmt.Println("\n ")
		testCmdFlagSet.Parse(os.Args[2:])
		runQueries(*testHostFlag,
			*testPortFlag,
			*testUserFlag,
			*testPassFlag,
			*testPathFlag,
			*testDatabaseFlag,
			*testDBTypeFlag,
			*testRepetitionFlag)
	case "list":
		fmt.Println("listing queries...")
		fmt.Println("\n ")
		listCmdFlagSet.Parse(os.Args[2:])
		list(*listPathFlag)
	default:
		fmt.Printf("no matching commands or arguments...")
	}
}

func runQueries(host string,
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

func testConnection(host string,
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

func list(path string) {
	queries, err := GetQueriesFromPath(path)

	if err != nil {
		fmt.Println("Erro buscando as querys: ", err)
	}

	for i, query := range queries {
		fmt.Println("Query num.: ", i)
		fmt.Println("Query: ", query.QueryString)
	}
}
