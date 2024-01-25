package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttp2RpcResponse Response Object
type DeleteHttp2RpcResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttp2RpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttp2RpcResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttp2RpcResponse", string(data)}, " ")
}
