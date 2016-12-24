package main

import (
	"os"

	"github.com/byakheee/yesman"
)

var (
	version  string
	revision string
)

func main() {
	os.Exit(yesman.Run(version, revision))
}
