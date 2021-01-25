package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateWhitelistRequestBody struct {
	Whitelist *UpdateWhitelistReq `json:"whitelist"`
}

func (o UpdateWhitelistRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateWhitelistRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistRequestBody", string(data)}, " ")
}
