package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateModuleDescriptionPrimitiveTypeHolder struct {

	// 私有模块（private-module）的描述。可用于客户识别被管理的私有模块。如果想要更新私有模块的描述，可以通过UpdatePrivateModuleMetadata API。
	ModuleDescription *string `json:"module_description,omitempty"`
}

func (o PrivateModuleDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateModuleDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateModuleDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
