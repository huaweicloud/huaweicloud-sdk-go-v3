package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPluginRequest Request Object
type ModifyPluginRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// 插件id
	PluginId string `json:"plugin_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`

	Body *WasmPlugin `json:"body,omitempty"`
}

func (o ModifyPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPluginRequest struct{}"
	}

	return strings.Join([]string{"ModifyPluginRequest", string(data)}, " ")
}
