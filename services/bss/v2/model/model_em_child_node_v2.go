package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EmChildNodeV2 struct {

	// 实体关系ID。
	RelationId *string `json:"relation_id,omitempty"`

	// 节点ID（即组织单元的Party ID）。
	Id *string `json:"id,omitempty"`

	// 节点名称。
	Name *string `json:"name,omitempty"`

	// 子节点列表。
	ChildNodes *[]EmChildNodeV2 `json:"child_nodes,omitempty"`
}

func (o EmChildNodeV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmChildNodeV2 struct{}"
	}

	return strings.Join([]string{"EmChildNodeV2", string(data)}, " ")
}
