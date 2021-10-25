package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBlockchainFlavorsRequest struct {
}

func (o ShowBlockchainFlavorsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlockchainFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainFlavorsRequest", string(data)}, " ")
}
