package exception

import (
	"fmt"
)

type AppValidate struct {
	Field string
	Tag   string
}

type AppError struct {
	Code           string
	StatusCode     uint32
	ErrorMessageTh string
	ErrorMessageEn string
	Fields         []AppValidate
}

func (e AppError) Error() string {
	return fmt.Sprintf("Code: %s, ErrorMessageTh: %s, ErrorMessageEn: %s", e.Code, e.ErrorMessageTh, e.ErrorMessageEn)
}
