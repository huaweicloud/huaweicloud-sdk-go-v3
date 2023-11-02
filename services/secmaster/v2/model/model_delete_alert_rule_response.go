package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRuleResponse Response Object
type DeleteAlertRuleResponse struct {

	// 告警规则 ID。Alert rule ID.
	RuleId *string `json:"rule_id,omitempty"`

	// 删除时间。Delete time.
	DeleteTime *int64 `json:"delete_time,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlertRuleResponse", string(data)}, " ")
}
