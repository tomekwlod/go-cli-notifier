package main

import (
	"flag"
	"fmt"
	"goclinotifier/client"
	"os"
)

var (
	backendURIFlag = flag.String("backend", "http://localhost:8080", "Backend API usage")
	helpFlag       = flag.Bool("help", false, "Help message")
)

func main() {
	flag.Parse()

	s := client.NewSwitch(*backendURIFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("command switch error: %s", err)
		os.Exit(2)
	}

}
