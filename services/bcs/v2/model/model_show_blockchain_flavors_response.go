package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowBlockchainFlavorsResponse struct {
	EnterpriseSpec *InstanceSpc `json:"enterprise_spec,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowBlockchainFlavorsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlockchainFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockchainFlavorsResponse", string(data)}, " ")
}
