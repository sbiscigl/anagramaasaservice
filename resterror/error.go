package resterror

import (
	"encoding/json"
	"fmt"
)

/*RestError type for my application*/
type RestError struct {
	Cause string `json:"cause"`
	Code  int    `json:"code"`
}

/*NewError returns a new error type*/
func NewError(cause string, code int) *RestError {
	return &RestError{
		Cause: cause,
		Code:  code,
	}
}

/*Error interface implimentatoin for error interface*/
func (er *RestError) Error() string {
	return er.Cause
}

/*ToJSON casts a error to a writable json*/
func (er *RestError) ToJSON() []byte {
	j, err := json.Marshal(er)
	if err != nil {
		fmt.Println("Json marshalling of a error failed")
	}
	return j
}
