package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeyAliasInfo struct {
	// 密钥ID。

	KeyId *string `json:"key_id,omitempty"`
	// 密钥别名。

	KeyAlias *string `json:"key_alias,omitempty"`
}

func (o KeyAliasInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyAliasInfo struct{}"
	}

	return strings.Join([]string{"KeyAliasInfo", string(data)}, " ")
}
