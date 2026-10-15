package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFederationsRequest Request Object
type ListFederationsRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：联邦端点名称，模糊查询。 **取值范围**：长度为1~63个字符。
	Name *string `json:"name,omitempty"`
}

func (o ListFederationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFederationsRequest struct{}"
	}

	return strings.Join([]string{"ListFederationsRequest", string(data)}, " ")
}
