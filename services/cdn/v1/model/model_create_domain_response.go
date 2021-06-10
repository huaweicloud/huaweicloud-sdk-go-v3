package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateDomainResponse struct {
	Domain         *CreateDomainResponseBodyContent `json:"domain,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CreateDomainResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDomainResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainResponse", string(data)}, " ")
}
