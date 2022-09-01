package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议配置信息
type RestConfConfigDto struct {

	// 是否需要发送会议邮件通知。默认值由企业级配置决定。 - True: 需要。 - False: 不需要。
	IsSendNotify *bool `json:"isSendNotify,omitempty" xml:"isSendNotify"`

	// 是否需要发送会议短信通知。默认值由企业级配置决定。 - True: 需要。 - False: 不需要。
	IsSendSms *bool `json:"isSendSms,omitempty" xml:"isSendSms"`

	// 是否需要发送会议通知。默认值由企业级配置决定。 - True: 需要。 - False: 不需要。
	IsSendCalendar *bool `json:"isSendCalendar,omitempty" xml:"isSendCalendar"`

	// 来宾入会,软终端是否自动静音。默认值由企业级配置决定。 - True: 自动静音。 - False: 不自动静音。
	IsAutoMute *bool `json:"isAutoMute,omitempty" xml:"isAutoMute"`

	// 来宾入会,硬终端是否自动静音。默认值由企业级配置决定。 - True: 自动静音。 - False: 不自动静音。
	IsHardTerminalAutoMute *bool `json:"isHardTerminalAutoMute,omitempty" xml:"isHardTerminalAutoMute"`

	// 是否来宾免密（仅随机会议有效）。 - True: 免密。 - False: 需要密码。
	IsGuestFreePwd *bool `json:"isGuestFreePwd,omitempty" xml:"isGuestFreePwd"`

	// 允许呼入的范围。 - 0: 所有用户。 - 2: 企业内用户。 - 3: 被邀请用户。
	CallInRestriction *int32 `json:"callInRestriction,omitempty" xml:"callInRestriction"`

	// 是否允许来宾启动会议(随机会议)。 - True: 允许来宾启动会议。 - False: 禁止来宾启动会议。
	AllowGuestStartConf *bool `json:"allowGuestStartConf,omitempty" xml:"allowGuestStartConf"`

	// 来宾密码
	GuestPwd *string `json:"guestPwd,omitempty" xml:"guestPwd"`

	// |参数名称：专用VMR会议ID类型 |参数描述：专用VMR会议ID类型 0: 固定ID 1: 随机ID |取值范围：[0,1]|
	VmrIDType *int32 `json:"vmrIDType,omitempty" xml:"vmrIDType"`

	// |参数名称：自动延长会议时长，0表示会议不延长 |建议取值范围：[0,60]|
	ProlongLength *int32 `json:"prolongLength,omitempty" xml:"prolongLength"`

	// 开启或者关闭等候室
	EnableWaitingRoom *bool `json:"enableWaitingRoom,omitempty" xml:"enableWaitingRoom"`
}

func (o RestConfConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestConfConfigDto struct{}"
	}

	return strings.Join([]string{"RestConfConfigDto", string(data)}, " ")
}
