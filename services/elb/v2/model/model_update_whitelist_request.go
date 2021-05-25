package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateWhitelistRequest struct {
	// 待更新的白名单id

	WhitelistId string `json:"whitelist_id"`

	Body *UpdateWhitelistRequestBody `json:"body,omitempty"`
}

func (o UpdateWhitelistRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateWhitelistRequest struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistRequest", string(data)}, " ")
}
