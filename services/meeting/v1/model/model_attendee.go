package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 与会者信息
type Attendee struct {

	// 与会者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// 与会者帐号，兼容终端老版本。如果没有携带userUUID，就通过accountId查询用户信息。
	AccountId *string `json:"accountId,omitempty"`

	// 与会者名称或昵称，长度限制为96个字符。
	Name string `json:"name"`

	// 会议中的角色。 - 0: 普通与会者。 - 1: 会议主席。 - 2: 预留字段，暂不对外开放。 default: 0
	Role *int32 `json:"role,omitempty"`

	// 电话号码(可支持SIP、TEL号码格式)。最大不超过127个字符。phone、email和sms三者需至少填写一个。当type为telepresence时，且设备为三屏智真，则该字段填写中屏号码。
	Phone string `json:"phone"`

	// 预留字段，取值类型同phone。当type为telepresence时，且设备为三屏智真，则该字段填写左屏号码
	Phone2 *string `json:"phone2,omitempty"`

	// 预留字段，取值类型同phone。当type为telepresence时，且设备为三屏智真，则该字段填写右屏号码
	Phone3 *string `json:"phone3,omitempty"`

	// 邮件地址。最大不超过255个字符。phone、email和sms三者需至少填写一个。
	Email *string `json:"email,omitempty"`

	// 短信通知的手机号码。最大不超过32个字符。phone、email和sms三者需至少填写一个。
	Sms *string `json:"sms,omitempty"`

	// 默认值由会议AS定义，号码类型枚举如下： - normal: 软终端。 - telepresence: 智真。单屏、三屏智真均属此类。（预留字段） - terminal: 会议室或硬终端。 - outside: 外部与会人。 - mobile: 用户手机号码。 - telephone: 软终端用户固定电话，暂不使用。
	Type string `json:"type"`

	// 部门ID。最大不超过64个字符。
	DeptUUID *string `json:"deptUUID,omitempty"`

	// 部门名称。最大不超过128个字符。
	DeptName *string `json:"deptName,omitempty"`
}

func (o Attendee) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attendee struct{}"
	}

	return strings.Join([]string{"Attendee", string(data)}, " ")
}
