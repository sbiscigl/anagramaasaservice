package error

import (
	"encoding/json"
	"fmt"
)

/*Error type for my application*/
type Error struct {
	Cause string `json:"cause"`
	Code  int    `json:"code"`
}

/*NewError returns a new error type*/
func NewError(cause string, code int) Error {
	return Error{
		Cause: cause,
		Code:  code,
	}
}

/*ToJSON casts a error to a writable json*/
func (er Error) ToJSON() []byte {
	j, err := json.Marshal(er)
	if err != nil {
		fmt.Println("Json marshalling of a error failed")
	}
	return j
}
