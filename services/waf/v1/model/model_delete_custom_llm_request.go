package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomLlmRequest Request Object
type DeleteCustomLlmRequest struct {

	// **参数解释：** 自定义大模型ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LlmId string `json:"llmId"`
}

func (o DeleteCustomLlmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomLlmRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomLlmRequest", string(data)}, " ")
}
