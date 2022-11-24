package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetachApiFromPluginRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 插件编号
	PluginId string `json:"plugin_id"`

	Body *PluginOperApiInfo `json:"body,omitempty"`
}

func (o DetachApiFromPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachApiFromPluginRequest struct{}"
	}

	return strings.Join([]string{"DetachApiFromPluginRequest", string(data)}, " ")
}
