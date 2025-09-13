package core

import (
    "fmt"
    "log"
)

func HandleDataError(err string) {
    errorMessage := fmt.Sprintf("\033[31mError: invalid data format, %s\033[0m\n", err)
    log.Fatal(errorMessage)
}

func HandleError(err string) {
	errorMessage := fmt.Sprintf("\033[31mError: %s\033[0m\n", err)
	log.Fatal(errorMessage)
}
