package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会场信息
type ParticipantInfo struct {

	// 与会者的号码。
	ParticipantID *string `json:"participantID,omitempty" xml:"participantID"`

	// 与会者的名称（昵称）。
	Name *string `json:"name,omitempty" xml:"name"`

	// 与会者的号码（预留字段）。
	SubscriberID *string `json:"subscriberID,omitempty" xml:"subscriberID"`

	// 会议中的角色。 - 1: 会议主持人。 - 0: 普通与会者。
	Role *int32 `json:"role,omitempty" xml:"role"`

	// 用户状态。目前固定返回MEETTING。
	State *string `json:"state,omitempty" xml:"state"`

	// 终端所在会议室信息。（预留字段）
	Address *string `json:"address,omitempty" xml:"address"`

	// - normal: 软终端。 - telepresence: 智真。单屏、三屏智真均属此类。（预留字段） - terminal: 会议室或硬终端。 - outside: 外部与会人。 - mobile: 用户手机号码。 - telephone: 用户固定电话。（预留字段）
	AttendeeType *string `json:"attendeeType,omitempty" xml:"attendeeType"`

	// 预订者的账号ID。
	AccountId *string `json:"accountId,omitempty" xml:"accountId"`

	// 当attendeeType为telepresence时，且设备为三屏智真，则该字段填写左屏号码。（预留字段）
	Phone2 *string `json:"phone2,omitempty" xml:"phone2"`

	// 当attendeeType为telepresence时，且设备为三屏智真，则该字段填写右屏号码。（预留字段）
	Phone3 *string `json:"phone3,omitempty" xml:"phone3"`

	// 邮件地址。最大不超过255个字符。
	Email *string `json:"email,omitempty" xml:"email"`

	// 短信通知的手机号码。最大不超过127个字符。
	Sms *string `json:"sms,omitempty" xml:"sms"`

	// 部门名称。最大不超过96个字符。
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 预订者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 第三方应用ID。
	AppId *string `json:"appId,omitempty" xml:"appId"`

	// 会议开始时是否自动邀请该与会者。
	IsAutoInvite *int32 `json:"isAutoInvite,omitempty" xml:"isAutoInvite"`

	// 是否不叠加会场名。
	IsNotOverlayPidName *bool `json:"isNotOverlayPidName,omitempty" xml:"isNotOverlayPidName"`
}

func (o ParticipantInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParticipantInfo struct{}"
	}

	return strings.Join([]string{"ParticipantInfo", string(data)}, " ")
}
