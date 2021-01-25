package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteConnctionRequest struct {
	ConnectionName string `json:"connection_name"`
}

func (o DeleteConnctionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteConnctionRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnctionRequest", string(data)}, " ")
}
