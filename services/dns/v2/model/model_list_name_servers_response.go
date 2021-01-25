package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListNameServersResponse struct {
	Nameservers    *[]NameServersResp `json:"nameservers,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListNameServersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNameServersResponse struct{}"
	}

	return strings.Join([]string{"ListNameServersResponse", string(data)}, " ")
}
