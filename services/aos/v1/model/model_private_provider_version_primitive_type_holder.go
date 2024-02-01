package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateProviderVersionPrimitiveTypeHolder struct {

	// provider的版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义
	ProviderVersion *string `json:"provider_version,omitempty"`
}

func (o PrivateProviderVersionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateProviderVersionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateProviderVersionPrimitiveTypeHolder", string(data)}, " ")
}
