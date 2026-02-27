package perror

import "fmt"

type Perror struct {
	Message string
	Code    int
}

func New(message string, code int) *Perror {
	return &Perror{
		Message: message,
		Code:    code,
	}
}

func (per *Perror) Error() string {
	return fmt.Sprintf("Error Message %v and Status %v", per.Message, per.Code)
}
