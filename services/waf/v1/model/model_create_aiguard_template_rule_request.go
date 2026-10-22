package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAiguardTemplateRuleRequest Request Object
type CreateAiguardTemplateRuleRequest struct {

	// **参数解释：** 策略ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 语言 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** en-us
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateAiguardTemplateRuleRequestBody `json:"body,omitempty"`
}

func (o CreateAiguardTemplateRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiguardTemplateRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateAiguardTemplateRuleRequest", string(data)}, " ")
}
