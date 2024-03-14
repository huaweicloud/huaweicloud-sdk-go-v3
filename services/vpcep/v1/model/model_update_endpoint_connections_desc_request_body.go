package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEndpointConnectionsDescRequestBody struct {

	// 连接管理描述字段列表
	Connections []ConnectionsDesc `json:"connections"`
}

func (o UpdateEndpointConnectionsDescRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointConnectionsDescRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointConnectionsDescRequestBody", string(data)}, " ")
}
