package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestScheduleConfDto 会议信息。
type RestScheduleConfDto struct {

	// 会议开始时间（UTC时间）。格式：yyyy-MM-dd HH:mm。 > * 创建预约会议时，如果没有指定开始时间或填空串，则表示会议马上开始 > * 时间是UTC时间，即0时区的时间
	StartTime *string `json:"startTime,omitempty"`

	// 会议持续时长，单位分钟。默认30分钟。最大1440分钟（24小时），最小15分钟。
	Length *int32 `json:"length,omitempty"`

	// 会议主题。最多128个字符。
	Subject *string `json:"subject,omitempty"`

	// 会议的媒体类型。 - Voice: 语音会议 - HDVideo: 视频会议
	MediaTypes string `json:"mediaTypes"`

	// 软终端创建即时会议时在当前字段带临时群组ID，由服务器在邀请其他与会者时在或者conference-info头域中携带。长度限制为31个字符。
	Groupuri *string `json:"groupuri,omitempty"`

	// 与会者列表。
	Attendees *[]RestAttendeeDto `json:"attendees,omitempty"`

	// 会议是否自动启动录制，在录播类型为:录播、直播+录播时有效。默认为不自动启动。 - 0: 不自动启动录制 - 1: 自动启动录制
	IsAutoRecord *int32 `json:"isAutoRecord,omitempty"`

	// 会议媒体加密模式。默认值由企业级的配置填充。 - 0: 自适应加密 - 1: 强制加密 - 2: 不加密
	EncryptMode *int32 `json:"encryptMode,omitempty"`

	// 会议通知短信或邮件的语言。默认中文。 - zh-CN: 简体中文 - en-US: 美国英文
	Language *string `json:"language,omitempty"`

	// 会议通知中会议时间的时区信息。时区信息，参考[[时区映射关系](https://support.huaweicloud.com/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hws)[[时区映射关系](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hk)。 > * 举例：“timeZoneID”:\"26\"，则通过华为云会议发送的会议通知中的时间将会标记为如“2021/11/11 星期四 00:00 - 02:00 (GMT) 格林威治标准时间:都柏林, 爱丁堡, 里斯本, 伦敦”。 > * 非周期会议，如果会议通知是通过第三方系统发送，则这个字段不用填写。
	TimeZoneID *string `json:"timeZoneID,omitempty"`

	// 录播类型。默认为禁用。 - 0: 禁用 - 1: 直播 - 2: 录播 - 3: 直播+录播
	RecordType *int32 `json:"recordType,omitempty"`

	// 主流直播推流地址，在录播类型为 :直播、直播+录播时有效。最大不超过255个字符。
	LiveAddress *string `json:"liveAddress,omitempty"`

	// 辅流直播推流地址，在录播类型为 :直播、直播+录播时有效。最大不超过255个字符。
	AuxAddress *string `json:"auxAddress,omitempty"`

	// 是否录制辅流，在录播类型为：录播、录播+直播时有效。默认只录制视频主流，不录制辅流。  - 0: 不录制  - 1: 录制
	RecordAuxStream *int32 `json:"recordAuxStream,omitempty"`

	ConfConfigInfo *RestConfConfigDto `json:"confConfigInfo,omitempty"`

	// 录播观看鉴权方式，在录播类型为:录播、直播+录播时有效。 - 0: 可通过链接观看/下载 - 1: 企业用户可观看/下载 - 2: 与会者可观看/下载
	RecordAuthType *int32 `json:"recordAuthType,omitempty"`

	// 是否使用云会议室或者个人会议ID召开预约会议。默认0。 - 0: 不使用云会议室或者个人会议ID - 1: 使用云会议室或者个人会议ID
	VmrFlag *int32 `json:"vmrFlag,omitempty"`

	CycleParams *CycleParams `json:"cycleParams,omitempty"`

	// 绑定给当前创会帐号的VMR ID。通过[[查询云会议室及个人会议ID](https://support.huaweicloud.com/api-meeting/meeting_21_1106.html)](tag:hws)[[查询云会议室及个人会议ID](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_1106.html)](tag:hk)接口获取。 > * vmrID取上述查询接口中返回的id，不是vmrId > * 创建个人会议ID的会议时，使用vmrMode=0的VMR；创建云会议室的会议时，使用vmrMode=1的VMR
	VmrID *string `json:"vmrID,omitempty"`

	// 会议最大与会人数。默认值0。 * 0：无限制 * 大于0：会议最大与会人数
	ConcurrentParticipants *int32 `json:"concurrentParticipants,omitempty"`
}

func (o RestScheduleConfDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestScheduleConfDto struct{}"
	}

	return strings.Join([]string{"RestScheduleConfDto", string(data)}, " ")
}
