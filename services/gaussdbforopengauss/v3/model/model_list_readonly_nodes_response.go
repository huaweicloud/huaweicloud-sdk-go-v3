package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReadonlyNodesResponse Response Object
type ListReadonlyNodesResponse struct {

	// **参数解释**: 只读节点列表。 **约束限制**: 不涉及。
	Nodes *[]ListReadonlyNodesResult `json:"nodes,omitempty"`

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	MaxReadonlyNodeNum *int32 `json:"max_readonly_node_num,omitempty"`
	HttpStatusCode     int    `json:"-"`
}

func (o ListReadonlyNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReadonlyNodesResponse struct{}"
	}

	return strings.Join([]string{"ListReadonlyNodesResponse", string(data)}, " ")
}
