package utils

import (
	log "log"
)

// PanicOnError : Prints the error & exits the program
func PanicOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s\n", msg, err)
	}
}
