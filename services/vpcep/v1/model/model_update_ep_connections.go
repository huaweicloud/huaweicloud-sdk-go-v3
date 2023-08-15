package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEpConnections struct {

	// 连接管理描述字段列表
	Connections []ConnectionsDesc `json:"connections"`
}

func (o UpdateEpConnections) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEpConnections struct{}"
	}

	return strings.Join([]string{"UpdateEpConnections", string(data)}, " ")
}
