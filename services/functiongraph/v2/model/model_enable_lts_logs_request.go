package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type EnableLtsLogsRequest struct {
}

func (o EnableLtsLogsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnableLtsLogsRequest struct{}"
	}

	return strings.Join([]string{"EnableLtsLogsRequest", string(data)}, " ")
}
