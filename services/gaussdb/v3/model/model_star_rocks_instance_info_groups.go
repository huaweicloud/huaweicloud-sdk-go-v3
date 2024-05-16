package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StarRocksInstanceInfoGroups struct {

	// 分组ID。
	Id *string `json:"id,omitempty"`

	// 分组名。
	Name *string `json:"name,omitempty"`

	// 实例节点。
	Nodes *[]StarRocksInstanceInfoNodes `json:"nodes,omitempty"`

	// 实例分组类型名。
	GroupTypeName *string `json:"group_type_name,omitempty"`

	// 实例分组节点类型。
	GroupNodeType *string `json:"group_node_type,omitempty"`
}

func (o StarRocksInstanceInfoGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoGroups struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoGroups", string(data)}, " ")
}
