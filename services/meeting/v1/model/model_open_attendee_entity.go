package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 与会嘉宾列表。该列表可以用于发送会议通知、会议提醒、会议开始时候进行自动邀请。
type OpenAttendeeEntity struct {
	// App ID，应用标识，一个应用只需创建一次。必填字段

	AppId string `json:"appId"`
	// 与会者的账号ID。如果是账号/密码鉴权场景，表示华为云会议帐号ID。如果是APP ID鉴权场景，表示第三方的User ID。必填字段

	UserAccount *string `json:"userAccount,omitempty"`
	// 与会者名称或昵称。长度限制为96个字符

	UserName *string `json:"userName,omitempty"`
	// 部门名称

	DeptName *string `json:"deptName,omitempty"`
	// 号码（可支持SIP、TEL号码格式）。最大不超过127个字符。phone、email和sms三者需至少填写一个。

	Phone *string `json:"phone,omitempty"`
	// 邮件地址。最大不超过255个字符。phone、email和sms三者需至少填写一个（用于预定、修改、取消会议的信息通知）。

	Email *string `json:"email,omitempty"`
	// 短信通知的手机号码。最大不超过32个字符。phone、email和sms三者需至少填写一个。（用于预定、修改、取消会议的信息通知）。

	Sms *string `json:"sms,omitempty"`
	// 是否硬终端（会议室或硬终端）。

	IsHardTerminal *bool `json:"isHardTerminal,omitempty"`
}

func (o OpenAttendeeEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenAttendeeEntity struct{}"
	}

	return strings.Join([]string{"OpenAttendeeEntity", string(data)}, " ")
}
