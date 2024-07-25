package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAlertRuleResponse Response Object
type EnableAlertRuleResponse struct {

	// Alert rule ID.
	FailList *[]AlertRule `json:"fail_list,omitempty"`

	// Alert rule ID.
	SuccessList *[]AlertRule `json:"success_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"EnableAlertRuleResponse", string(data)}, " ")
}
