package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteFirewallResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFirewallResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFirewallResponse struct{}"
	}

	return strings.Join([]string{"DeleteFirewallResponse", string(data)}, " ")
}
