package cmd

import (
	"flag"
	"fmt"
	"os"
)

func ParseArgs() {
	testCmdFlagSet := flag.NewFlagSet("test", flag.ExitOnError)
	// testHostFlag := testCmdFlagSet.String("host", "localhost", "the database host")
	// testPortFlag := testCmdFlagSet.Int("port", 42069, "the database port")
	// testUserFlag := testCmdFlagSet.String("user", "root", "the database user")
	// testPassFlag := testCmdFlagSet.String("pass", "password", "the database password")
	// testPathFlag := testCmdFlagSet.String("path", "queries", "the queries file or directory")

	listCmdFlagSet := flag.NewFlagSet("list", flag.ExitOnError)
	// listPathFlag := listCmdFlagSet.String("path", "queries", "the queries file or directory")

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "test":
		fmt.Println("running tests...")
		testCmdFlagSet.Parse(os.Args[2:])
	case "list":
		fmt.Println("listing queries...")
		listCmdFlagSet.Parse(os.Args[2:])
	default:
		fmt.Printf("no matching commands or arguments...")
	}
}
