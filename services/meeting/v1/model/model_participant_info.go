package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会场信息
type ParticipantInfo struct {

	// 与会者的号码。
	ParticipantID *string `json:"participantID,omitempty"`

	// 与会者的名称（昵称）。
	Name *string `json:"name,omitempty"`

	// 与会者的号码（预留字段）。
	SubscriberID *string `json:"subscriberID,omitempty"`

	// 会议中的角色。 - 1: 会议主持人。 - 0: 普通与会者。
	Role *int32 `json:"role,omitempty"`

	// 用户状态。目前固定返回MEETTING。
	State *string `json:"state,omitempty"`

	// 终端所在会议室信息。（预留字段）
	Address *string `json:"address,omitempty"`

	// - normal: 软终端。 - telepresence: 智真。单屏、三屏智真均属此类。（预留字段） - terminal: 会议室或硬终端。 - outside: 外部与会人。 - mobile: 用户手机号码。 - telephone: 用户固定电话。（预留字段）
	AttendeeType *string `json:"attendeeType,omitempty"`

	// 预订者的账号ID。
	AccountId *string `json:"accountId,omitempty"`

	// 当attendeeType为telepresence时，且设备为三屏智真，则该字段填写左屏号码。（预留字段）
	Phone2 *string `json:"phone2,omitempty"`

	// 当attendeeType为telepresence时，且设备为三屏智真，则该字段填写右屏号码。（预留字段）
	Phone3 *string `json:"phone3,omitempty"`

	// 邮件地址。最大不超过255个字符。
	Email *string `json:"email,omitempty"`

	// 短信通知的手机号码。最大不超过127个字符。
	Sms *string `json:"sms,omitempty"`

	// 部门名称。最大不超过96个字符。
	DeptName *string `json:"deptName,omitempty"`

	// 预订者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// 第三方应用ID。
	AppId *string `json:"appId,omitempty"`

	// 会议开始时是否自动邀请该与会者。
	IsAutoInvite *int32 `json:"isAutoInvite,omitempty"`

	// 是否不叠加会场名。
	IsNotOverlayPidName *bool `json:"isNotOverlayPidName,omitempty"`
}

func (o ParticipantInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParticipantInfo struct{}"
	}

	return strings.Join([]string{"ParticipantInfo", string(data)}, " ")
}
