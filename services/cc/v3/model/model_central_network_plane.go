package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkPlane 中心网络平面详情信息。
type CentralNetworkPlane struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 关联的中心网络ER实例列表。
	AssociateErTables *[]AssociateErTableDocument `json:"associate_er_tables,omitempty"`

	// 当自动连接所有ER实例时，排除中心网络的ER实例的连接。
	ExcludeErConnections *[][]AssociateErInstanceDocument `json:"exclude_er_connections,omitempty"`
}

func (o CentralNetworkPlane) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkPlane struct{}"
	}

	return strings.Join([]string{"CentralNetworkPlane", string(data)}, " ")
}
