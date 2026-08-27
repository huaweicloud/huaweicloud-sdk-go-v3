package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDrainPoolNodesReq 批量对节点进行排水的请求体。
type BatchDrainPoolNodesReq struct {

	// **参数解释**：节点名称列表。 **约束限制**：节点不能超过10个。
	NodeNames []string `json:"nodeNames"`
}

func (o BatchDrainPoolNodesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDrainPoolNodesReq struct{}"
	}

	return strings.Join([]string{"BatchDrainPoolNodesReq", string(data)}, " ")
}
