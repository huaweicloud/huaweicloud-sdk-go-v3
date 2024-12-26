package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestAttendeeDto 与会者信息。
type RestAttendeeDto struct {

	// 与会者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// 与会者的帐号。 * 如果是帐号/密码鉴权场景：选填，表示华为云会议帐号 * 如果是APPID鉴权场景：必填，表示第三方的User ID，同时需要携带参数appId
	AccountId *string `json:"accountId,omitempty"`

	// 与会者名称。长度限制为96个字符。
	Name *string `json:"name,omitempty"`

	// 会议中的角色。默认为普通与会者。 - 0: 普通与会者 - 1: 会议主持人
	Role *int32 `json:"role,omitempty"`

	// 号码。支持SIP号码或者手机号码。 * 如果是帐号/密码鉴权场景：必填 * 如果是APP ID鉴权场景：选填 > * 号码可以通过[[查询企业通讯](https://support.huaweicloud.com/api-meeting/meeting_21_0512.html)](tag:hws)[[查询企业通讯](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0512.html)](tag:hk)接口录获取。返回的number是SIP号码，phone是手机号码 > * 填SIP号码系统会呼叫对应的软终端或者硬终端；填手机号码系统会呼叫手机 > * 呼叫手机需要开通PSTN权限，否则无法呼叫
	Phone *string `json:"phone,omitempty"`

	// 预留字段，取值类型同phone。
	Phone2 *string `json:"phone2,omitempty"`

	// 预留字段，取值类型同phone。
	Phone3 *string `json:"phone3,omitempty"`

	// 邮箱地址。需要发邮件通知时填写。
	Email *string `json:"email,omitempty"`

	// 短信通知的手机号码。需要发短信通知时填写。
	Sms *string `json:"sms,omitempty"`

	// 用户入会时是否需要自动静音。默认不静音。 - 0: 不需要静音 - 1: 需要静音 > 仅会中邀请与会者时生效。
	IsMute *int32 `json:"isMute,omitempty"`

	// 会议开始时是否自动邀请该与会者。默认值由企业级配置决定。 - 0: 不自动邀请 - 1: 自动邀请 > 仅并发会议资源的随机会议ID会议才生效。
	IsAutoInvite *int32 `json:"isAutoInvite,omitempty"`

	// 终端类型，类型枚举如下： * normal: 软终端 * terminal: 会议室或硬终端 * outside: 外部与会人 * mobile: 用户手机号码 * ideahub：ideahub * board: 电子白板（SmartRooms），含Maxhub、海信大屏、IdeaHub B2 * hwvision：华为智慧屏TV
	Type *string `json:"type,omitempty"`

	// 预留字段，终端所在会议室信息。
	Address *string `json:"address,omitempty"`

	// 部门ID。
	DeptUUID *string `json:"deptUUID,omitempty"`

	// 部门名称。最大不超过128个字符。
	DeptName *string `json:"deptName,omitempty"`

	// App ID。如果是APP ID鉴权场景，此项必填。参考[[App ID的申请](https://support.huaweicloud.com/devg-meeting/meeting_20_0011.html#section1)](tag:hws)[[App ID的申请](https://support.huaweicloud.com/intl/zh-cn/devg-meeting/meeting_20_0011.html#section1)](tag:hk)。
	AppId *string `json:"appId,omitempty"`

	// 企业内唯一会场标识, 0标识为普通与会者，1标识为企业内唯一会场; uniqueType 为1， 同时type要指定为customnumber
	UniqueType *int32 `json:"uniqueType,omitempty"`
}

func (o RestAttendeeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestAttendeeDto struct{}"
	}

	return strings.Join([]string{"RestAttendeeDto", string(data)}, " ")
}
