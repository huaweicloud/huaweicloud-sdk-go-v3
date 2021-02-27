package model

import (
	"encoding/json"

	"strings"
)

//
type IdpToken struct {
	// token产生时间。
	IssuedAt string `json:"issued_at"`
	// token到期时间。
	ExpiresAt string `json:"expires_at"`
	// 获取token的方式。
	Methods []string           `json:"methods"`
	User    *UnscopedTokenUser `json:"user"`
}

func (o IdpToken) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IdpToken struct{}"
	}

	return strings.Join([]string{"IdpToken", string(data)}, " ")
}