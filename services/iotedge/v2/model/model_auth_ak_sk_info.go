package model

import (
	"encoding/json"

	"strings"
)

type AuthAkSkInfo struct {
	// 鉴权秘钥

	Secret *string `json:"secret,omitempty"`
}

func (o AuthAkSkInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AuthAkSkInfo struct{}"
	}

	return strings.Join([]string{"AuthAkSkInfo", string(data)}, " ")
}
