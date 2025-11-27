package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateModuleVersionPrimitiveTypeHolder struct {

	// 模块的版本号。版本号必须遵循语义化版本号（Semantic Version），为用户自定义
	ModuleVersion *string `json:"module_version,omitempty"`
}

func (o PrivateModuleVersionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateModuleVersionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateModuleVersionPrimitiveTypeHolder", string(data)}, " ")
}
