package model

import (
	"encoding/json"

	"strings"
)

type AccessRomaBriefInfo struct {
	// 认证key，加密存储

	AppKey *string `json:"app_key,omitempty"`
}

func (o AccessRomaBriefInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AccessRomaBriefInfo struct{}"
	}

	return strings.Join([]string{"AccessRomaBriefInfo", string(data)}, " ")
}
