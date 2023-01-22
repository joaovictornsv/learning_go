package mail_validator

import (
	"fmt"

	"github.com/badoux/checkmail"
	"github.com/joaovictornsv/learning_go/packages/logger"
)

func Validate(email string) {
	erro := checkmail.ValidateFormat(email)
	if erro == nil {
		logger.Log("Email ok")
	} else {
		fmt.Println(erro)
	}
}
