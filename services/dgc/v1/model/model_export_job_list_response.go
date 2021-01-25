package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExportJobListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportJobListResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportJobListResponse struct{}"
	}

	return strings.Join([]string{"ExportJobListResponse", string(data)}, " ")
}
