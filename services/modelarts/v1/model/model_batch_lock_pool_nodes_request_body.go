package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchLockPoolNodesRequestBody struct {

	// **参数解释**：需要变更锁状态的节点名称列表。 **约束限制**：不涉及。
	NodeNames []string `json:"nodeNames"`

	// **参数解释**：变更的功能类型。 **约束限制**：不涉及。
	Actions []string `json:"actions"`
}

func (o BatchLockPoolNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchLockPoolNodesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchLockPoolNodesRequestBody", string(data)}, " ")
}
