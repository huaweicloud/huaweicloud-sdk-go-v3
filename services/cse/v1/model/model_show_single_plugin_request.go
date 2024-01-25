package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSinglePluginRequest Request Object
type ShowSinglePluginRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// 插件id
	PluginId string `json:"plugin_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`
}

func (o ShowSinglePluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinglePluginRequest struct{}"
	}

	return strings.Join([]string{"ShowSinglePluginRequest", string(data)}, " ")
}
