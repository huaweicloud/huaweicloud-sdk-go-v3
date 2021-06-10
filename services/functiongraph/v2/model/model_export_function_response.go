package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExportFunctionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportFunctionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportFunctionResponse struct{}"
	}

	return strings.Join([]string{"ExportFunctionResponse", string(data)}, " ")
}
