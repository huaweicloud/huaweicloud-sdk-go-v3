package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttp2RpcRequest Request Object
type DeleteHttp2RpcRequest struct {

	// 网关实例id
	GatewayId string `json:"gateway_id"`

	// 插件id
	Http2RpcId string `json:"http2Rpc_id"`

	// 该字段内容填为 \"application/json\"
	Accept *string `json:"Accept,omitempty"`
}

func (o DeleteHttp2RpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttp2RpcRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttp2RpcRequest", string(data)}, " ")
}
