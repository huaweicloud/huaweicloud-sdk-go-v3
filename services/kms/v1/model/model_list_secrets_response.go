package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSecretsResponse struct {
	// 凭据详情列表。

	Secrets *[]Secret `json:"secrets,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecretsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSecretsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretsResponse", string(data)}, " ")
}
