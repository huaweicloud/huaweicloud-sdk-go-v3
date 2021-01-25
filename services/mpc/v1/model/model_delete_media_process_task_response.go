package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteMediaProcessTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMediaProcessTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMediaProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteMediaProcessTaskResponse", string(data)}, " ")
}
