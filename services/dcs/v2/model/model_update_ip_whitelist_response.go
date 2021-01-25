package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateIpWhitelistResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateIpWhitelistResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateIpWhitelistResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpWhitelistResponse", string(data)}, " ")
}
