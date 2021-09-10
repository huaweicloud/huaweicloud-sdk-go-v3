package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisableDomainResponse struct {
	Domain         *DomainsWithPort `json:"domain,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DisableDomainResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisableDomainResponse struct{}"
	}

	return strings.Join([]string{"DisableDomainResponse", string(data)}, " ")
}
