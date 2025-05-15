package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkPlaneDocument 中心网络平面文档。
type CentralNetworkPlaneDocument struct {

	// 实例名称。
	Name string `json:"name"`

	// 关联的中心网络ER实例列表。
	AssociateErTables *[]AssociateErTableDocument `json:"associate_er_tables,omitempty"`

	// 当自动连接所有ER实例时，排除中心网络的ER实例的连接。
	ExcludeErConnections *[][]AssociateErInstanceDocument `json:"exclude_er_connections,omitempty"`
}

func (o CentralNetworkPlaneDocument) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkPlaneDocument struct{}"
	}

	return strings.Join([]string{"CentralNetworkPlaneDocument", string(data)}, " ")
}
