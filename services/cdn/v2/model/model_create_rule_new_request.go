package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleNewRequest Request Object
type CreateRuleNewRequest struct {

	// **参数解释：** 加速域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	Body *CreateRuleRequest `json:"body,omitempty"`
}

func (o CreateRuleNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleNewRequest struct{}"
	}

	return strings.Join([]string{"CreateRuleNewRequest", string(data)}, " ")
}
