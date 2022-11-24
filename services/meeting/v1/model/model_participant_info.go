package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 被邀请的与会者信息。包含预约会议时邀请的与会者和会中主持人邀请的与会者。
type ParticipantInfo struct {

	// 与会者的号码。
	ParticipantID *string `json:"participantID,omitempty"`

	// 与会者的名称。
	Name *string `json:"name,omitempty"`

	// 与会者的号码（预留字段）。
	SubscriberID *string `json:"subscriberID,omitempty"`

	// 与会者的角色。 - 1: 会议主持人 - 0: 普通与会者
	Role *int32 `json:"role,omitempty"`

	// 用户状态。目前固定返回MEETTING。
	State *string `json:"state,omitempty"`

	// 终端所在会议室信息（预留字段）。
	Address *string `json:"address,omitempty"`

	// 与会者终端类型。 - normal: 软终端。 - terminal: 会议室或硬终端。 - outside: 外部与会人。 - mobile: 用户手机号码。
	AttendeeType *string `json:"attendeeType,omitempty"`

	// 预订者的帐号。 * 如果是帐号/密码鉴权场景，表示华为云会议帐号 * 如果是APP ID鉴权场景，表示第三方的User ID
	AccountId *string `json:"accountId,omitempty"`

	// 预留字段。
	Phone2 *string `json:"phone2,omitempty"`

	// 预留字段。
	Phone3 *string `json:"phone3,omitempty"`

	// 邮箱地址。
	Email *string `json:"email,omitempty"`

	// 短信通知的手机号码。
	Sms *string `json:"sms,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`

	// 预订者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// App ID。参考[[App ID的申请](https://support.huaweicloud.com/devg-meeting/meeting_20_0011.html#section1)](tag:hws)[[App ID的申请](https://support.huaweicloud.com/intl/zh-cn/devg-meeting/meeting_20_0011.html#section1)](tag:hk)。
	AppId *string `json:"appId,omitempty"`

	// 会议开始时是否自动邀请该与会者。默认值由企业级配置决定。 * 0： 不自动邀请 * 1： 自动邀请
	IsAutoInvite *int32 `json:"isAutoInvite,omitempty"`

	// 是否不叠加会场名（VDC场景下适用）。
	IsNotOverlayPidName *bool `json:"isNotOverlayPidName,omitempty"`
}

func (o ParticipantInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParticipantInfo struct{}"
	}

	return strings.Join([]string{"ParticipantInfo", string(data)}, " ")
}
