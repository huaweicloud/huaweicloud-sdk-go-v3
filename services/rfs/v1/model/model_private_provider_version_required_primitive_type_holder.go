package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateProviderVersionRequiredPrimitiveTypeHolder struct {

	// provider的版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义
	ProviderVersion string `json:"provider_version"`
}

func (o PrivateProviderVersionRequiredPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateProviderVersionRequiredPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateProviderVersionRequiredPrimitiveTypeHolder", string(data)}, " ")
}
