package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListServicesResponse struct {
	// 服务信息列表。

	Services *[]Service `json:"services,omitempty"`

	Links          *Links `json:"links,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o KeystoneListServicesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListServicesResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListServicesResponse", string(data)}, " ")
}
