package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClickHouseNodeInfoResponseBodyNodeList struct {

	// ClickHouse实例节点id。
	NodeId string `json:"node_id"`

	// ClickHouse实例节点名称。
	NodeName string `json:"node_name"`

	// 节点主备角色。
	Role string `json:"role"`
}

func (o ClickHouseNodeInfoResponseBodyNodeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseNodeInfoResponseBodyNodeList struct{}"
	}

	return strings.Join([]string{"ClickHouseNodeInfoResponseBodyNodeList", string(data)}, " ")
}
