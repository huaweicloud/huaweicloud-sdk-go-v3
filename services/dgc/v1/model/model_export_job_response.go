package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExportJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportJobResponse struct{}"
	}

	return strings.Join([]string{"ExportJobResponse", string(data)}, " ")
}
