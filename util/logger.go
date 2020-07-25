package util

import (
	"log"
	"os"
)

var Logger *log.Logger = log.New(os.Stdout, "=> ", 0)
