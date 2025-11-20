package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDetailfNodesInfo nodes参数说明。
type GetDetailfNodesInfo struct {

	// DDM实例节点状态。
	Status string `json:"status"`

	// DDM实例节点port。
	Port string `json:"port"`

	// DDM实例节点IP。
	Ip string `json:"ip"`

	// 节点所在组ID。
	GroupId string `json:"group_id"`

	// 节点ID。
	NodeId string `json:"node_id"`
}

func (o GetDetailfNodesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDetailfNodesInfo struct{}"
	}

	return strings.Join([]string{"GetDetailfNodesInfo", string(data)}, " ")
}
