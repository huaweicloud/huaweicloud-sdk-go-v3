package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2ComputeRequest Request Object
type ShowV2ComputeRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：计算资源ID。可通过[查询计算资源列表]获取计算资源ID(ListV2Computes.xml)。 **约束限制**：不涉及。 **取值范围**：长度为1-36个英文字母、数字和中划线的组合。 **默认取值**：不涉及。
	ComputeId string `json:"compute_id"`
}

func (o ShowV2ComputeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2ComputeRequest struct{}"
	}

	return strings.Join([]string{"ShowV2ComputeRequest", string(data)}, " ")
}
