package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeNodeInfo struct {

	// **参数解释**：节点批次ID，批次变更时需要，可以从节点的os.modelarts.node/batch.uid标签中获取。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	BatchUID *string `json:"batchUID,omitempty"`

	// **参数解释**：批次缩容场景，指定要缩容的节点名称列表。 **约束限制**：不涉及。
	DeleteNodeNames *[]string `json:"deleteNodeNames,omitempty"`
}

func (o ResizeNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeNodeInfo struct{}"
	}

	return strings.Join([]string{"ResizeNodeInfo", string(data)}, " ")
}
