package utils

import "log"

func HandleError(err error) {
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}
}
