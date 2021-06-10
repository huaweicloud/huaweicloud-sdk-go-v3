package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainDetailResponse struct {
	Domain         *Domains `json:"domain,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowDomainDetailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainDetailResponse", string(data)}, " ")
}
