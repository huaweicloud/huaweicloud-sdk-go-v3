package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateProviderVersionDescriptionPrimitiveTypeHolder struct {

	// 私有provider版本（provider version）的描述。可用于客户识别并管理私有provider的版本。注意：provider版本为不可更新（immutable），所以该字段不可更新，如果需要更新，请删除后重建
	VersionDescription *string `json:"version_description,omitempty"`
}

func (o PrivateProviderVersionDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateProviderVersionDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateProviderVersionDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
