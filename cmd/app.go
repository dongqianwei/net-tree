package main

import "log"
import "os"

var logger = log.New(os.Stdout, "[INFO] ", log.Lshortfile)

func main() {
	logger.Println("agent start...")
}