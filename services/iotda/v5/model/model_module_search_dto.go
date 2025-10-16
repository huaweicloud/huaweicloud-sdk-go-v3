package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModuleSearchDto OTA模块结构。
type ModuleSearchDto struct {

	// 模块名称。
	ModuleName string `json:"module_name"`

	// 模块版本号。
	ModuleVersion string `json:"module_version"`
}

func (o ModuleSearchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleSearchDto struct{}"
	}

	return strings.Join([]string{"ModuleSearchDto", string(data)}, " ")
}
