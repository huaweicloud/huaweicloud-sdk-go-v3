package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTranscodingTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTranscodingTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTranscodingTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingTaskResponse", string(data)}, " ")
}
