package main

import (
	"os"
)
func init() {
	os.Stdout.Write([]byte("validating mariadb"))
}

func main() {
	os.Exit(0)
}
