package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 加密数据项配置
type EncryptDataItem struct {

	// 加密数据项键名
	Name string `json:"name"`

	// 加密数据项键值
	Value string `json:"value"`

	// 加密数据项键值是否已加密，默认为false
	IsEncrypted *bool `json:"is_encrypted,omitempty"`
}

func (o EncryptDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataItem struct{}"
	}

	return strings.Join([]string{"EncryptDataItem", string(data)}, " ")
}
