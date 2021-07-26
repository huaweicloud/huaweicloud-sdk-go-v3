package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePremiumHostRequest struct {
	Body *CreatePremiumHostRequestBody `json:"body,omitempty"`
}

func (o CreatePremiumHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePremiumHostRequest struct{}"
	}

	return strings.Join([]string{"CreatePremiumHostRequest", string(data)}, " ")
}
