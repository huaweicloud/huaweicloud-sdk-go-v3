package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenScheduleConfReq 创建网络研讨会请求。
type OpenScheduleConfReq struct {

	// 网络研讨会主题。长度限制为128个字符。
	Subject string `json:"subject"`

	// 网络研讨会描述，长度限制为1000个字符。
	Description *string `json:"description,omitempty"`

	// 网络研讨会开始时间（UTC时间），格式“yyyy-MM-dd HH:mm”。
	StartTime string `json:"startTime"`

	// 网络研讨会持续时长，单位分钟，取值范围[15,1440]。
	Duration int32 `json:"duration"`

	// 会议通知中会议时间的时区信息。时区信息，参考[[时区映射关系](https://support.huaweicloud.com/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hws)[[时区映射关系](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hk)。 > * 举例：“timeZoneID”:\"26\"，则通过华为云会议发送的会议通知中的时间将会标记为如“2021/11/11 星期四 00:00 - 02:00 (GMT) 格林威治标准时间:都柏林, 爱丁堡, 里斯本, 伦敦”。
	TimeZoneId int32 `json:"timeZoneId"`

	// 与会嘉宾列表。 > 观众只能自己通过链接或者会议ID+密码加入，不支持被邀请。
	Attendees *[]OpenAttendeeEntity `json:"attendees,omitempty"`

	NotifySetting *OpenNotifySetting `json:"notifySetting,omitempty"`

	// 绑定给当前创建网络研讨会帐号的VMR ID。通过[[查询云会议室及个人会议ID](https://support.huaweicloud.com/api-meeting/meeting_21_1106.html)](tag:hws)[[查询云会议室及个人会议ID](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_1106.html)](tag:hk)接口获取。 > * vmrID取上述查询接口中返回的id，不是vmrId。 > * 创建网络研讨会时，使用vmrMode=2的VMR。
	VmrID string `json:"vmrID"`

	// 嘉宾密码（4-16位长度的纯数字)。不指定则系统自动生成。
	GuestPasswd *string `json:"guestPasswd,omitempty"`

	// 观众密码（4-16位长度的纯数字)。不指定则系统自动生成。
	AudiencePasswd *string `json:"audiencePasswd,omitempty"`

	// 入会范围开关。 * false：不限制入会范围 * true：限制入会范围
	CallRestriction *bool `json:"callRestriction,omitempty"`

	// 主持人、嘉宾入会范围 。仅callRestriction设置成true后生效。 * 0: 所有用户 * 2: 企业内用户 * 3: 被邀请用户
	Scope *int32 `json:"scope,omitempty"`

	// 观众入会范围。仅callRestriction设置成true后生效。 * 0: 所有用户 * 2: 企业内用户和被邀请用户
	AudienceScope *int32 `json:"audienceScope,omitempty"`

	EnableRecording *YesNoEnum `json:"enableRecording,omitempty"`
}

func (o OpenScheduleConfReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenScheduleConfReq struct{}"
	}

	return strings.Join([]string{"OpenScheduleConfReq", string(data)}, " ")
}
