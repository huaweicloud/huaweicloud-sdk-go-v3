package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableAlertRuleResponse Response Object
type DisableAlertRuleResponse struct {

	// Alert rule ID.
	FailList *[]AlertRule `json:"fail_list,omitempty"`

	// Alert rule ID.
	SuccessList *[]AlertRule `json:"success_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"DisableAlertRuleResponse", string(data)}, " ")
}
