package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateWhitelistRequestBody struct {
	Whitelist *CreateWhitelistReq `json:"whitelist"`
}

func (o CreateWhitelistRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWhitelistRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWhitelistRequestBody", string(data)}, " ")
}
