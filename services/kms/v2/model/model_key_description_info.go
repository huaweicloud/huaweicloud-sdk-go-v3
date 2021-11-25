package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeyDescriptionInfo struct {
	// 密钥ID。

	KeyId *string `json:"key_id,omitempty"`
	// 密钥描述。

	KeyDescription *string `json:"key_description,omitempty"`
}

func (o KeyDescriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyDescriptionInfo struct{}"
	}

	return strings.Join([]string{"KeyDescriptionInfo", string(data)}, " ")
}
