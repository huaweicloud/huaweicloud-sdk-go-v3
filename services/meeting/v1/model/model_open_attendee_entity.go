package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenAttendeeEntity 与会嘉宾列表。
type OpenAttendeeEntity struct {

	// App ID。如果是APP ID鉴权场景，此项必填。参考[[App ID的申请](https://support.huaweicloud.com/devg-meeting/meeting_20_0011.html#section1)](tag:hws)[[App ID的申请](https://support.huaweicloud.com/intl/zh-cn/devg-meeting/meeting_20_0011.html#section1)](tag:hk)。
	AppId string `json:"appId"`

	// 嘉宾的帐号。 * 如果是帐号/密码鉴权场景: 选填，表示华为云会议帐号ID * 如果是APP ID鉴权场景：必填，表示第三方的User ID，同时需要携带参数appId
	UserAccount *string `json:"userAccount,omitempty"`

	// 嘉宾的名称。长度限制为96个字符。
	UserName *string `json:"userName,omitempty"`

	// 部门名称，最大128字符。
	DeptName *string `json:"deptName,omitempty"`

	// 号码。支持SIP号码或者手机号码。 * 如果是帐号/密码鉴权场景：必填 * 如果是APP ID鉴权场景：选填 > * 号码可以通过[[查询企业通讯](https://support.huaweicloud.com/api-meeting/meeting_21_0512.html)](tag:hws)[[查询企业通讯](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0512.html)](tag:hk)接口录获取。返回的number是SIP号码，phone是手机号码 > * 填SIP号码系统会呼叫对应的软终端或者硬终端；填手机号码系统会呼叫手机 > * 呼叫手机需要开通PSTN权限，否则无法呼叫
	Phone *string `json:"phone,omitempty"`

	// 邮件地址。需要发邮件通知时填写。
	Email *string `json:"email,omitempty"`

	// 短信通知的手机号码。需要发短信通知时填写。
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
