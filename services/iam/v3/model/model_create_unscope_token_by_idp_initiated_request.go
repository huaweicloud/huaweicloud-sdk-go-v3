package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateUnscopeTokenByIdpInitiatedRequest struct {
	XIdpId string       `json:"X-Idp-Id"`
	Body   *interface{} `json:"body,omitempty"`
}

func (o CreateUnscopeTokenByIdpInitiatedRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUnscopeTokenByIdpInitiatedRequest struct{}"
	}

	return strings.Join([]string{"CreateUnscopeTokenByIdpInitiatedRequest", string(data)}, " ")
}
