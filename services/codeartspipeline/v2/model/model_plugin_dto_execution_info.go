package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginDtoExecutionInfo 执行信息
type PluginDtoExecutionInfo struct {

	// 执行信息
	InnerExecutionInfo *interface{} `json:"inner_execution_info,omitempty"`
}

func (o PluginDtoExecutionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDtoExecutionInfo struct{}"
	}

	return strings.Join([]string{"PluginDtoExecutionInfo", string(data)}, " ")
}
