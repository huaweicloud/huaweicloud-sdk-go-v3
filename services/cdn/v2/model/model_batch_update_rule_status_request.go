package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateRuleStatusRequest Request Object
type BatchUpdateRuleStatusRequest struct {

	// **参数解释：** 加速域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	Body *BatchUpdateRulesRequest `json:"body,omitempty"`
}

func (o BatchUpdateRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateRuleStatusRequest", string(data)}, " ")
}
