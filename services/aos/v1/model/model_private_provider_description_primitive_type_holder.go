package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateProviderDescriptionPrimitiveTypeHolder struct {

	// 私有provider（private-provider）的描述。可用于客户识别被管理的私有provider。
	ProviderDescription *string `json:"provider_description,omitempty"`
}

func (o PrivateProviderDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateProviderDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateProviderDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
