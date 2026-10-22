package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomLlmRequest Request Object
type ListCustomLlmRequest struct {

	// **参数解释：** 自定义大模型名称筛选 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 模型名称筛选 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ModelName *string `json:"model_name,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 每页数量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCustomLlmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomLlmRequest struct{}"
	}

	return strings.Join([]string{"ListCustomLlmRequest", string(data)}, " ")
}
