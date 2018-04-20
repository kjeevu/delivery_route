package main

import (
	"flag"

	"delivery_route/tracking"
)

var fileName = flag.String("file_name", "points.csv", "Accepting the file name from command Line")

func main() {
	flag.Parse()
	tracking.ProcessInfo(*fileName)
}
