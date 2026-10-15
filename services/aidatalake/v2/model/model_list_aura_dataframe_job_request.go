package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuraDataframeJobRequest Request Object
type ListAuraDataframeJobRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：端点ID。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：作业ID。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：分页数量。 **约束限制**：不涉及。 **取值范围**：长度为1~100。 **默认取值**：不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页偏移。 **约束限制**：不涉及。 **取值范围**：长度为0~100。 **默认取值**：不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：作业状态。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`
}

func (o ListAuraDataframeJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuraDataframeJobRequest struct{}"
	}

	return strings.Join([]string{"ListAuraDataframeJobRequest", string(data)}, " ")
}
