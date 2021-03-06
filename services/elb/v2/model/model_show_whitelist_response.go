package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowWhitelistResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowWhitelistResponse struct{}"
	}

	return strings.Join([]string{"ShowWhitelistResponse", string(data)}, " ")
}
