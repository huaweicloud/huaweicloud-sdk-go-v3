package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 与会者信息
type RestAttendeeDto struct {

	// 与会者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// 与会者的账号ID。 如果是账号/密码鉴权场景，选填，表示华为云会议帐号ID。 如果是APPID鉴权场景，必填，表示第三方的User ID，同时需要携带appid参数。
	AccountId *string `json:"accountId,omitempty"`

	// 与会者名称或昵称，长度限制为96个字符。
	Name *string `json:"name,omitempty"`

	// 会议中的角色。默认为普通与会者。 - 0: 普通与会者。 - 1: 会议主席。 - 2: 预留字段，暂不对外开放。
	Role *int32 `json:"role,omitempty"`

	// 如果是账号/密码鉴权场景，必填，号码（可支持SIP、TEL号码格式）。 如果是APP ID鉴权场景，选填。 最大不超过127个字符。phone、email和sms三者需至少填写一个。
	Phone *string `json:"phone,omitempty"`

	// 预留字段，取值类型同phone。
	Phone2 *string `json:"phone2,omitempty"`

	// 预留字段，取值类型同phone。
	Phone3 *string `json:"phone3,omitempty"`

	// 邮件地址。最大不超过255个字符。phone、email和sms三者需至少填写一个。
	Email *string `json:"email,omitempty"`

	// 短信通知的手机号码。最大不超过32个字符。phone、email和sms三者需至少填写一个。
	Sms *string `json:"sms,omitempty"`

	// 用户入会时是否需要自动静音。默认不静音。 - 0: 不需要静音。 - 1: 需要静音。
	IsMute *int32 `json:"isMute,omitempty"`

	// 会议开始时是否自动邀请该与会者。默认值由企业级配置决定。 - 0: 不自动邀请 - 1: 自动邀请
	IsAutoInvite *int32 `json:"isAutoInvite,omitempty"`

	// 默认值由会议AS定义，号码类型枚举如下： - normal: 软终端。 - telepresence: 智真。单屏、三屏智真均属此类。（预留字段） - terminal: 会议室或硬终端。 - outside: 外部与会人。 - mobile: 用户手机号码。 - telephone: 软终端用户固定电话，暂不使用。
	Type *string `json:"type,omitempty"`

	// 终端所在会议室信息。（预留字段）
	Address *string `json:"address,omitempty"`

	// 部门ID。最大不超过64个字符。
	DeptUUID *string `json:"deptUUID,omitempty"`

	// 部门名称。最大不超过128个字符。
	DeptName *string `json:"deptName,omitempty"`

	// App ID，应用标识，一个应用只需创建一次。如果是APP ID鉴权场景，此项必填。
	AppId *string `json:"appId,omitempty"`
}

func (o RestAttendeeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestAttendeeDto struct{}"
	}

	return strings.Join([]string{"RestAttendeeDto", string(data)}, " ")
}
