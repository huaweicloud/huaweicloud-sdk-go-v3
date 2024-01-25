package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginsRequest Request Object
type ShowPluginsRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`
}

func (o ShowPluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginsRequest struct{}"
	}

	return strings.Join([]string{"ShowPluginsRequest", string(data)}, " ")
}
