package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议其他配置信息。
type RestConfConfigDto struct {

	// 是否需要发送会议邮件通知。默认值由企业级配置决定。 - true: 需要 - false: 不需要
	IsSendNotify *bool `json:"isSendNotify,omitempty"`

	// 是否需要发送会议短信通知。默认值由企业级配置决定。 - true: 需要 - false: 不需要 > 只有正式商用企业才有发送会议短信通知的权限，免费企业即使isSendSms设置成true也不会发送会议短信通知。
	IsSendSms *bool `json:"isSendSms,omitempty"`

	// 是否需要发送会议邮件日历通知。默认值由企业级配置决定。 - true: 需要 - false: 不需要
	IsSendCalendar *bool `json:"isSendCalendar,omitempty"`

	// 来宾入会,软终端是否自动静音。默认值由企业级配置决定。 - true: 自动静音 - false: 不自动静音
	IsAutoMute *bool `json:"isAutoMute,omitempty"`

	// 来宾入会,硬终端是否自动静音。默认值由企业级配置决定。 - true: 自动静音 - false: 不自动静音
	IsHardTerminalAutoMute *bool `json:"isHardTerminalAutoMute,omitempty"`

	// 是否来宾免密。 - true: 免密 - false: 需要密码 > 仅随机会议ID的会议生效。
	IsGuestFreePwd *bool `json:"isGuestFreePwd,omitempty"`

	// 允许加入会议的范围。 - 0: 所有用户 - 2: 企业内用户 - 3: 被邀请用户
	CallInRestriction *int32 `json:"callInRestriction,omitempty"`

	// 是否允许来宾启动会议。 - true: 允许来宾启动会议 - false: 禁止来宾启动会议 > 仅随机会议ID的会议生效。
	AllowGuestStartConf *bool `json:"allowGuestStartConf,omitempty"`

	// 来宾密码（4-16位长度的纯数字）。
	GuestPwd *string `json:"guestPwd,omitempty"`

	// 云会议室的会议ID模式。 * 0：固定会议ID * 1：随机会议ID
	VmrIDType *int32 `json:"vmrIDType,omitempty"`

	// 自动延长会议时长（取值范围0-60）。 * 0：表示会议到点自动结束，不延长会议 * 其他：表示自动延长的时长 > * 自动结束会议是按照会议时长计算。比如预定的会议是9:00开始11:00结束，会议时长2个小时，如果与会者8:00就加入会议了，那会议在10:00就会自动结束 > * 设置成其他值时，只要会中还有与会者，会议自动多次延长
	ProlongLength *int32 `json:"prolongLength,omitempty"`

	// 开启或者关闭等候室。 * true：开启 * false：不开启
	EnableWaitingRoom *bool `json:"enableWaitingRoom,omitempty"`
}

func (o RestConfConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestConfConfigDto struct{}"
	}

	return strings.Join([]string{"RestConfConfigDto", string(data)}, " ")
}
