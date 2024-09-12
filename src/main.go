package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: timestamp <timestamp>")
		return
	}

	timestampStr := os.Args[1]
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		fmt.Println("Error: invalid timestamp")
		return
	}

	tm := time.Unix(timestamp, 0).UTC()
	fmt.Println("Date and Time (UTC):", tm.Format(time.RFC3339))
}
