package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ValidateDictionaryResponse struct {
	// 字典名称 - 字符集：中文、英文字母、数字、下划线和空格 - 约束：实例下唯一

	Name *string `json:"name,omitempty"`
	// 字典编码 - 字符集：英文字母、数字、下划线和空格 - 约束：实例下唯一

	Code           *string `json:"code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateDictionaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDictionaryResponse struct{}"
	}

	return strings.Join([]string{"ValidateDictionaryResponse", string(data)}, " ")
}
