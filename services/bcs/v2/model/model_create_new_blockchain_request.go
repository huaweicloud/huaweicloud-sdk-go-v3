package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateNewBlockchainRequest struct {
	Body *CreateRequestBody `json:"body,omitempty"`
}

func (o CreateNewBlockchainRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateNewBlockchainRequest struct{}"
	}

	return strings.Join([]string{"CreateNewBlockchainRequest", string(data)}, " ")
}
