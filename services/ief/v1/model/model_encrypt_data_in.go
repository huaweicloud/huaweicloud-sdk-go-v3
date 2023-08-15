package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncryptDataIn 加密数据信息
type EncryptDataIn struct {

	// 加密数据名称，小写英文字母、数字、中划线，以小写字母或数字开头，最大长度为64个字符，不能为空
	Name string `json:"name"`

	// 加密数据描述
	Description *string `json:"description,omitempty"`

	// 加密数据项配置
	Config []EncryptDataItem `json:"config"`
}

func (o EncryptDataIn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataIn struct{}"
	}

	return strings.Join([]string{"EncryptDataIn", string(data)}, " ")
}
