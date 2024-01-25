package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePluginRequest Request Object
type DeletePluginRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// 插件id
	PluginId string `json:"plugin_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`
}

func (o DeletePluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePluginRequest struct{}"
	}

	return strings.Join([]string{"DeletePluginRequest", string(data)}, " ")
}
