package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListActiveActiveDomainsResponse struct {
	// 双活域列表信息。

	Domains        *[]ShowActiveActiveDomainParams `json:"domains,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListActiveActiveDomainsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListActiveActiveDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListActiveActiveDomainsResponse", string(data)}, " ")
}
