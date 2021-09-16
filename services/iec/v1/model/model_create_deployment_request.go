package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDeploymentRequest struct {
	Body *CreateDeploymentRequestBody `json:"body,omitempty"`
}

func (o CreateDeploymentRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDeploymentRequest struct{}"
	}

	return strings.Join([]string{"CreateDeploymentRequest", string(data)}, " ")
}
