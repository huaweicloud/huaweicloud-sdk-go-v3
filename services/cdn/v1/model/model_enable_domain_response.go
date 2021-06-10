package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type EnableDomainResponse struct {
	Domain         *Domains `json:"domain,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o EnableDomainResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnableDomainResponse struct{}"
	}

	return strings.Join([]string{"EnableDomainResponse", string(data)}, " ")
}
