package model

import (
	"encoding/json"

	"strings"
)

type VaultBindRules struct {
	// 按tags过滤自动绑定的资源

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o VaultBindRules) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultBindRules struct{}"
	}

	return strings.Join([]string{"VaultBindRules", string(data)}, " ")
}
