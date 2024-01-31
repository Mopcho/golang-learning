package main

import (
	"fmt"
)

type userError struct {
	statusCode int32
	message    string
}

func (e userError) Error() string {
	return fmt.Sprintf("status code: %v; message: %v", e.statusCode, e.message)
}

func sendSMS(message string) (string, error) {
	if len(message) < 5 {
		return message, nil
	} else {
		return "", userError{message: "Message must be less than 5 length", statusCode: 400}
	}
}

func main() {
	msg := "123456"

	message, err := sendSMS(msg)

	fmt.Printf("message: %v | err: %v", message, err)
}
