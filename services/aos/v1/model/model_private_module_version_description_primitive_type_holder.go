package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateModuleVersionDescriptionPrimitiveTypeHolder struct {

	// 模块版本（module version）的描述。可用于客户识别并管理模块的版本。注意：模块版本为不可更新（immutable），即描述不可更新，如果需要更新，请删除后重建
	VersionDescription *string `json:"version_description,omitempty"`
}

func (o PrivateModuleVersionDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateModuleVersionDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateModuleVersionDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
