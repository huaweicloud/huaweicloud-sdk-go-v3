package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateWhitelistResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWhitelistResponse struct{}"
	}

	return strings.Join([]string{"CreateWhitelistResponse", string(data)}, " ")
}
