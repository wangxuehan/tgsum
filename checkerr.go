package main

import "log"

func checkErr(info string, err error) {
	if err != nil {
		log.Fatalf("%s %q:", info, err)
	}
}
