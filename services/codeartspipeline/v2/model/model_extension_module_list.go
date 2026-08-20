package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionModuleList 插件的modules的具体信息集合
type ExtensionModuleList struct {

	// 模块信息
	Data *[]ExtensionModule `json:"data,omitempty"`

	// modules集合长度
	Total *int32 `json:"total,omitempty"`
}

func (o ExtensionModuleList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionModuleList struct{}"
	}

	return strings.Join([]string{"ExtensionModuleList", string(data)}, " ")
}
