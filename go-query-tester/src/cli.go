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

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "test":
		fmt.Println("running tests...")
		testCmdFlagSet.Parse(os.Args[2:])
		RunTests(*testHostFlag,
			*testPortFlag,
			*testUserFlag,
			*testPassFlag,
			*testPathFlag,
			*testDatabaseFlag,
			*testDBTypeFlag,
			*testRepetitionFlag)
	case "list":
		fmt.Println("listing queries...")
		listCmdFlagSet.Parse(os.Args[2:])
		List(*listPathFlag)
	default:
		fmt.Printf("no matching commands or arguments...")
	}
}
