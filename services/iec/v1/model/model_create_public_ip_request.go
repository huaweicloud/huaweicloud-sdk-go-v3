package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePublicIpRequest struct {
	Body *CreatePublicIpRequestBody `json:"body,omitempty"`
}

func (o CreatePublicIpRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePublicIpRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicIpRequest", string(data)}, " ")
}
