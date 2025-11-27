package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateModuleVersionRequiredPrimitiveTypeHolder struct {

	// 模块的版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义
	ModuleVersion string `json:"module_version"`
}

func (o PrivateModuleVersionRequiredPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateModuleVersionRequiredPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateModuleVersionRequiredPrimitiveTypeHolder", string(data)}, " ")
}
