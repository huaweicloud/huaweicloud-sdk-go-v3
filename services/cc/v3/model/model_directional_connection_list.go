package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DirectionalConnectionList struct {

	// 有向连接列表。
	DirectionalConnections []DirectionalConnection `json:"directional_connections"`
}

func (o DirectionalConnectionList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectionalConnectionList struct{}"
	}

	return strings.Join([]string{"DirectionalConnectionList", string(data)}, " ")
}
