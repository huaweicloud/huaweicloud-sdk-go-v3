package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 加密数据项配置
type EncryptDataItem struct {

	// 加密数据项键名，小写英文字母、数字、中划线，以小写字母或数字开头，最大长度为32个字符，不能为空
	Name string `json:"name"`

	// 加密数据项键值
	Value string `json:"value"`

	// 加密数据项键值是否已加密，默认为true
	IsEncrypted *bool `json:"is_encrypted,omitempty"`
}

func (o EncryptDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataItem struct{}"
	}

	return strings.Join([]string{"EncryptDataItem", string(data)}, " ")
}
