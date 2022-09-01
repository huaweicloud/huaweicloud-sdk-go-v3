package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 密钥描述信息。
type KeyDescriptionInfo struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// 密钥描述。
	KeyDescription *string `json:"key_description,omitempty" xml:"key_description"`
}

func (o KeyDescriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyDescriptionInfo struct{}"
	}

	return strings.Join([]string{"KeyDescriptionInfo", string(data)}, " ")
}
