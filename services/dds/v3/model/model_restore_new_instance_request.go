package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreNewInstanceRequest struct {
	Body *RestoreNewInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreNewInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceRequest", string(data)}, " ")
}
