package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainRequest struct {
	Domain *string `json:"domain,omitempty"`
}

func (o ShowDomainRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainRequest", string(data)}, " ")
}
