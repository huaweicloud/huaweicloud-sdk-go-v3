package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainItemLocationDetailsResponse struct {
	// 域名详情数据列表

	Domains        *[]DomainItemLocationDetails `json:"domains,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowDomainItemLocationDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainItemLocationDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainItemLocationDetailsResponse", string(data)}, " ")
}
