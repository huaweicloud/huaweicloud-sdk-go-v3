package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRuleResponse Response Object
type DeleteAlertRuleResponse struct {

	// 是否删除.
	Deleted *bool `json:"deleted,omitempty"`

	// Alert rule ID.
	FailList *[]AlertRule `json:"fail_list,omitempty"`

	// Alert rule ID.
	SuccessList *[]AlertRule `json:"success_list,omitempty"`

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
