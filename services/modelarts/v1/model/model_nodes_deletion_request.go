package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodesDeletionRequest 批量删除节点请求体。
type NodesDeletionRequest struct {

	// **参数解释**：待删除的节点名称列表。 **约束限制**：不涉及。
	DeleteNodeNames []string `json:"deleteNodeNames"`
}

func (o NodesDeletionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodesDeletionRequest struct{}"
	}

	return strings.Join([]string{"NodesDeletionRequest", string(data)}, " ")
}
