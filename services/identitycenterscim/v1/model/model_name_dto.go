package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NameDto struct {

	// 包含要显示的名称的格式化版本的字符串
	Formatted *string `json:"formatted,omitempty"`

	// 用户的姓氏
	FamilyName string `json:"familyName"`

	// 用户的名字
	GivenName string `json:"givenName"`

	// 用户的中间名
	MiddleName *string `json:"middleName,omitempty"`

	// 用户的尊称前缀
	HonorificPrefix *string `json:"honorificPrefix,omitempty"`

	// 用户的尊称后缀
	HonorificSuffix *string `json:"honorificSuffix,omitempty"`
}

func (o NameDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameDto struct{}"
	}

	return strings.Join([]string{"NameDto", string(data)}, " ")
}
