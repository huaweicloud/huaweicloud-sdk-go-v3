package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2RayFederationsRequest Request Object
type ListV2RayFederationsRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：联邦端点名称，模糊查询。 **约束限制**：不涉及。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ListV2RayFederationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2RayFederationsRequest struct{}"
	}

	return strings.Join([]string{"ListV2RayFederationsRequest", string(data)}, " ")
}
