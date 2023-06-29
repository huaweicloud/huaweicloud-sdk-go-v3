package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Attendee 与会者信息。
type Attendee struct {

	// 与会者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// 与会者的华为云会议帐号。
	AccountId *string `json:"accountId,omitempty"`

	// 与会者名称，长度限制为96个字符。
	Name string `json:"name"`

	// 会议中的角色。默认为普通与会者。 - 0: 普通与会者 - 1: 会议主持人
	Role *int32 `json:"role,omitempty"`

	// 号码。支持SIP号码或者手机号码。 > * 号码可以通过[[查询企业通讯](https://support.huaweicloud.com/api-meeting/meeting_21_0512.html)](tag:hws)[[查询企业通讯](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0512.html)](tag:hk)接口录获取。返回的number是SIP号码，phone是手机号码 > * 填SIP号码系统会呼叫对应的软终端或者硬终端；填手机号码系统会呼叫手机 > * 呼叫手机需要开通PSTN权限，否则无法呼叫
	Phone string `json:"phone"`

	// 预留字段，取值类型同参数“phone”。
	Phone2 *string `json:"phone2,omitempty"`

	// 预留字段，取值类型同参数“phone”。
	Phone3 *string `json:"phone3,omitempty"`

	// 邮件地址。 > 会中邀请不发会议通知，不用填写。
	Email *string `json:"email,omitempty"`

	// 短信通知的手机号码。 > 会中邀请不发会议通知，不用填写。
	Sms *string `json:"sms,omitempty"`

	// 终端类型，类型枚举如下： * normal：软终端 * terminal：硬终端 * outside：外部与会人 * mobile：用户手机号码 * ideahub：ideahub * board: 电子白板（SmartRooms）。含Maxhub、海信大屏、IdeaHub B2hwvision：华为智慧屏TV
	Type string `json:"type"`

	// 部门编码。
	DeptUUID *string `json:"deptUUID,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`
}

func (o Attendee) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attendee struct{}"
	}

	return strings.Join([]string{"Attendee", string(data)}, " ")
}
