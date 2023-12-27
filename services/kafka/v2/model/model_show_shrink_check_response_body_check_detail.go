package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShrinkCheckResponseBodyCheckDetail broker检查结果
type ShowShrinkCheckResponseBodyCheckDetail struct {

	// broker序号
	BrokerId *int32 `json:"broker_id,omitempty"`

	// 节点是否可删除。
	CanDelete *bool `json:"can_delete,omitempty"`

	// 节点是否为zk部署节点。
	IsZkNode *bool `json:"is_zk_node,omitempty"`

	// broker是否为controller。
	IsController *bool `json:"is_controller,omitempty"`

	// broker上是否存在topic数据。
	HasTopics *bool `json:"has_topics,omitempty"`

	// broker上存在的topic列表。
	Topics *[]string `json:"topics,omitempty"`
}

func (o ShowShrinkCheckResponseBodyCheckDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShrinkCheckResponseBodyCheckDetail struct{}"
	}

	return strings.Join([]string{"ShowShrinkCheckResponseBodyCheckDetail", string(data)}, " ")
}
