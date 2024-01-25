package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttp2RpcRequest Request Object
type CreateHttp2RpcRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`

	Body *Http2Rpc `json:"body,omitempty"`
}

func (o CreateHttp2RpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttp2RpcRequest struct{}"
	}

	return strings.Join([]string{"CreateHttp2RpcRequest", string(data)}, " ")
}
