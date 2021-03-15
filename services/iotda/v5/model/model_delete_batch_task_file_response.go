package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteBatchTaskFileResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBatchTaskFileResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBatchTaskFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteBatchTaskFileResponse", string(data)}, " ")
}
