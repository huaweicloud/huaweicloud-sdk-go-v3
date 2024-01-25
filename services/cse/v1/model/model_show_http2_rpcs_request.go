package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttp2RpcsRequest Request Object
type ShowHttp2RpcsRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`
}

func (o ShowHttp2RpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttp2RpcsRequest struct{}"
	}

	return strings.Join([]string{"ShowHttp2RpcsRequest", string(data)}, " ")
}
