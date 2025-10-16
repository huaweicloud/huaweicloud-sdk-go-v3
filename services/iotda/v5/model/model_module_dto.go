package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModuleDto OTA模块结构。
type ModuleDto struct {

	// 模块名称。
	ModuleName string `json:"module_name"`

	// 模块版本号。
	ModuleVersion string `json:"module_version"`
}

func (o ModuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleDto struct{}"
	}

	return strings.Join([]string{"ModuleDto", string(data)}, " ")
}
