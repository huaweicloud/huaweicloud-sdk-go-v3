package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateWhitelistResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateWhitelistResponse struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistResponse", string(data)}, " ")
}
