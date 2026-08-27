package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDrainPoolNodesRequest Request Object
type BatchDrainPoolNodesRequest struct {

	// **参数解释**：资源池ID。取值资源池详情的metadata.name字段。 **约束限制**：不涉及。 **取值范围**：只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为36-63个字符。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	Body *BatchDrainPoolNodesReq `json:"body,omitempty"`
}

func (o BatchDrainPoolNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDrainPoolNodesRequest struct{}"
	}

	return strings.Join([]string{"BatchDrainPoolNodesRequest", string(data)}, " ")
}
