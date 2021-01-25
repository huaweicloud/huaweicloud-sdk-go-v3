package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteExtractTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExtractTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteExtractTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteExtractTaskResponse", string(data)}, " ")
}
