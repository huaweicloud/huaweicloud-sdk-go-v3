package model

import (
	"encoding/json"

	"strings"
)

type AccessRomaInfo struct {
	// 认证key，加密存储

	AppKey *string `json:"app_key,omitempty"`
	// 认证secret，加密存储

	AppSecret *string `json:"app_secret,omitempty"`
}

func (o AccessRomaInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AccessRomaInfo struct{}"
	}

	return strings.Join([]string{"AccessRomaInfo", string(data)}, " ")
}
