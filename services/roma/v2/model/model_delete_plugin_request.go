package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePluginRequest Request Object
type DeletePluginRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 插件编号
	PluginId string `json:"plugin_id"`
}

func (o DeletePluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePluginRequest struct{}"
	}

	return strings.Join([]string{"DeletePluginRequest", string(data)}, " ")
}
