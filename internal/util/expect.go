package util

import "log"

func IsOK(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
