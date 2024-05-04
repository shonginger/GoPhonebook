package logger

import (
	"log"
	"os"
)

// Info/Debug writes logs in the color blue with "INFO: " as prefix
var Info = func(err error) {
	log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile).Println(err.Error())
}

// Warning writes logs in the color yellow with "WARNING: " as prefix
var Warning = func(err error) {
	log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile).Println(err.Error())
}

// Error writes logs in the color red with "ERROR: " as prefix
var Error = func(err error) {
	log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile).Println(err.Error())
}
