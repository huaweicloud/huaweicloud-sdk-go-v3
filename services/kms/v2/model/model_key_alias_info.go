package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 密钥别名信息。
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
