package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDevServerJobTemplateRequest Request Object
type GetDevServerJobTemplateRequest struct {

	// **参数解释**：Lite Server任务模板id。 **约束限制**：1 - 64字符，字母、数字和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o GetDevServerJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevServerJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"GetDevServerJobTemplateRequest", string(data)}, " ")
}
