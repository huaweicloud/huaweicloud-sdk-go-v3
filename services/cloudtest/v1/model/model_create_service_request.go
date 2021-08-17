package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateServiceRequest struct {
	Body *ServiceRequestBody `json:"body,omitempty"`
}

func (o CreateServiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateServiceRequest struct{}"
	}

	return strings.Join([]string{"CreateServiceRequest", string(data)}, " ")
}
