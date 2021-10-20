package app

import (
	"log"

	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Printf("[info]", err.Key, " ", err.Message)
	}
	return
}
