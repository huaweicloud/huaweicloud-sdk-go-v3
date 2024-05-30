package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApproverParam struct {

	// 调用审核系统的应用名称，开发人员自己定。
	AppName *string `json:"app_name,omitempty"`

	// 审批人姓名。
	ApproverName string `json:"approver_name"`

	// 审批人ID。
	UserId string `json:"user_id"`

	// 邮箱地址。
	Email *string `json:"email,omitempty"`

	// 电话号码。
	PhoneNumber *string `json:"phone_number,omitempty"`

	// 邮件通知。
	EmailNotify *bool `json:"email_notify,omitempty"`

	// 短信通知。
	SmsNotify *bool `json:"sms_notify,omitempty"`
}

func (o ApproverParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproverParam struct{}"
	}

	return strings.Join([]string{"ApproverParam", string(data)}, " ")
}
