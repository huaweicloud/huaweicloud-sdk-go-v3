package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NameDto The full name of the user.
type NameDto struct {

	// 用户的姓氏
	FamilyName string `json:"family_name"`

	// 包含要显示的名称的格式化版本的字符串
	Formatted *string `json:"formatted,omitempty"`

	// 用户的名字
	GivenName string `json:"given_name"`

	// 用户的尊称前缀
	HonorificPrefix *string `json:"honorific_prefix,omitempty"`

	// 用户的尊称后缀
	HonorificSuffix *string `json:"honorific_suffix,omitempty"`

	// 用户的中间名
	MiddleName *string `json:"middle_name,omitempty"`
}

func (o NameDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameDto struct{}"
	}

	return strings.Join([]string{"NameDto", string(data)}, " ")
}
