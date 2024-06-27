package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChInstancesInfoRsponseInstanceGroups struct {

	// 分组ID。
	Id string `json:"id"`

	// 分组名。
	Name string `json:"name"`

	// 实例分组类型名，现在只支持clickhouse。
	GroupTypeName string `json:"group_type_name"`

	// 实例节点信息。
	Nodes []ClickHouseNodeInfo `json:"nodes"`
}

func (o ChInstancesInfoRsponseInstanceGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChInstancesInfoRsponseInstanceGroups struct{}"
	}

	return strings.Join([]string{"ChInstancesInfoRsponseInstanceGroups", string(data)}, " ")
}
