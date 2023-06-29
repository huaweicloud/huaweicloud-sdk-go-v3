package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePluginRequest Request Object
type UpdatePluginRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 插件编号
	PluginId string `json:"plugin_id"`

	Body *PluginCreate `json:"body,omitempty"`
}

func (o UpdatePluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginRequest struct{}"
	}

	return strings.Join([]string{"UpdatePluginRequest", string(data)}, " ")
}
