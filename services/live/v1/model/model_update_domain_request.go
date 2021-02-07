package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainRequest struct {
	Body *LiveDomainModifyReq `json:"body,omitempty"`
}

func (o UpdateDomainRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainRequest", string(data)}, " ")
}
