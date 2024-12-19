package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProviderResponseBody 运营商的返回体。
type ProviderResponseBody struct {

	// 运营商id
	Id *string `json:"id,omitempty"`

	// 运营商简写
	ShortName *string `json:"short_name,omitempty"`

	// 运营商类型
	Type *string `json:"type,omitempty"`

	ProviderValue *ProviderValueBody `json:"provider_value,omitempty"`

	// 运营商描述信息
	Description *string `json:"description,omitempty"`
}

func (o ProviderResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProviderResponseBody struct{}"
	}

	return strings.Join([]string{"ProviderResponseBody", string(data)}, " ")
}
