package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryConnectionsResponse struct {
	// 连接到该实例或节点的客户端IP地址。

	ClientIp string `json:"client_ip"`
	// 该IP对应的连接数。

	Count int32 `json:"count"`
}

func (o QueryConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryConnectionsResponse struct{}"
	}

	return strings.Join([]string{"QueryConnectionsResponse", string(data)}, " ")
}
