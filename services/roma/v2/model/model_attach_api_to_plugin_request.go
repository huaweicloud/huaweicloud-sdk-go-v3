package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachApiToPluginRequest Request Object
type AttachApiToPluginRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 插件编号
	PluginId string `json:"plugin_id"`

	Body *PluginOperApiInfo `json:"body,omitempty"`
}

func (o AttachApiToPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachApiToPluginRequest struct{}"
	}

	return strings.Join([]string{"AttachApiToPluginRequest", string(data)}, " ")
}
