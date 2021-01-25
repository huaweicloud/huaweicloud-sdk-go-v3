package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteBlockchainResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBlockchainResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBlockchainResponse struct{}"
	}

	return strings.Join([]string{"DeleteBlockchainResponse", string(data)}, " ")
}
