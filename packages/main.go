package main

import (
	mail_validator "github.com/joaovictornsv/learning_go/packages/check_mail"
	"github.com/joaovictornsv/learning_go/packages/logger"
)

func main() {
	logger.Log("Hello")

	mail_validator.Validate("hello@email.com")
}
