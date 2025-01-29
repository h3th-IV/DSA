package main

import (
	"errors"
	"log"
)

func Funcer(arg1 int, arg2 string) (ret1 int, ret2 string, ret3 error) {
	err := errors.New("a test error")
	return arg1, arg2, err
}

func ErrorHandler(arg1 int, arg2 string) {
	ret1, ret2, err := Funcer(arg1, arg2)
	if err != nil {
		log.Printf("an error occured")
	}
	log.Printf("returned value 1 is %v, retunred value 2 is %v", ret1, ret2)
}
