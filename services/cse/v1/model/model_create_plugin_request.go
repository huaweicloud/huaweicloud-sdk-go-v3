package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginRequest Request Object
type CreatePluginRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`

	Body *WasmPlugin `json:"body,omitempty"`
}

func (o CreatePluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginRequest struct{}"
	}

	return strings.Join([]string{"CreatePluginRequest", string(data)}, " ")
}
