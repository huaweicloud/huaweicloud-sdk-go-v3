package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type EnableLtsLogsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableLtsLogsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnableLtsLogsResponse struct{}"
	}

	return strings.Join([]string{"EnableLtsLogsResponse", string(data)}, " ")
}
