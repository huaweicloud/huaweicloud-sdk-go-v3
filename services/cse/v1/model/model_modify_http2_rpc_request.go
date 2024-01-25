package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyHttp2RpcRequest Request Object
type ModifyHttp2RpcRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// http2Rpc id
	Http2RpcId string `json:"http2Rpc_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`

	Body *Http2Rpc `json:"body,omitempty"`
}

func (o ModifyHttp2RpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHttp2RpcRequest struct{}"
	}

	return strings.Join([]string{"ModifyHttp2RpcRequest", string(data)}, " ")
}
