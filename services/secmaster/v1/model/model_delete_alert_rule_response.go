package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRuleResponse Response Object
type DeleteAlertRuleResponse struct {

	// 是否删除
	Deleted *bool `json:"deleted,omitempty"`

	// 失败列表
	FailList *[]AlertRule `json:"fail_list,omitempty"`

	// 成功列表
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
