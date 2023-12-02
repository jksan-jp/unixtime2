package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter a number")
		os.Exit(1)
	}

	arg := os.Args[1]
	// Try to parse the argument as an integer
	_, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		fmt.Println("Please enter a number")
	@	os.Exit(1)
	}

	// Get the Unix timestamp from the command line argument
	unixTimeStr := os.Args[1]
	unixTime, err := strconv.ParseInt(unixTimeStr, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the Unix timestamp to a time.Time object
	t := time.Unix(unixTime, 0)

	// Format the time.Time object to a string
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
