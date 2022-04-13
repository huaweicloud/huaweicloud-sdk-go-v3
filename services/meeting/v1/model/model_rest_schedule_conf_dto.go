package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建或修改会议请求体。
type RestScheduleConfDto struct {
	// 会议开始时间（UTC时间）。 预定创建会议时，如果没有指定开始时间，或填空串，则表示会议马上开始。 格式：yyyy-MM-dd HH:mm

	StartTime *string `json:"startTime,omitempty"`
	// 会议持续时长，单位分钟，最大值为1440，最短15。default：30。

	Length *int32 `json:"length,omitempty"`
	// 会议主题。长度限制为128个字符。

	Subject *string `json:"subject,omitempty"`
	// 会议的媒体类型。 由1个或多个枚举String组成，多个枚举时，每个枚举值之间通过”,”逗号分隔，枚举值如下： - Voice: 语音。 - Video: 标清视频。 - HDVideo: 高清视频（与Video互斥，如果同时选择Video、HDVideo，则系统默认选择Video） - Telepresence: 智真(与HDVideo、Video互斥，如果同时选择，系统使用Telepresence)。（预留字段） - Data: 多媒体（系统配置决定是否自动添加Data）。

	MediaTypes string `json:"mediaTypes"`
	// 软终端创建即时会议时在当前字段带临时群组ID，由服务器在邀请其他与会者时在或者conference-info头域中携带。长度限制为31个字符。

	Groupuri *string `json:"groupuri,omitempty"`
	// 与会者列表。该列表可以用于发送会议通知、会议提醒、会议开始时候进行自动邀请。

	Attendees *[]RestAttendeeDto `json:"attendees,omitempty"`
	// 会议是否自动启动录制，在录播类型为:录播、直播+录播时有效。默认为不自动启动。 - 0: 不自动启动录制 - 1: 自动启动录制

	IsAutoRecord *int32 `json:"isAutoRecord,omitempty"`
	// 会议媒体加密模式。默认值由企业级的配置填充。 - 0: 自适应加密 - 1: 强制加密 - 2: 不加密

	EncryptMode *int32 `json:"encryptMode,omitempty"`
	// 会议的默认语言，默认值由会议云服务定义。 对于系统支持的语言，按照RFC3066规范传递。 - zh-CN: 简体中文。 - en-US: 美国英文。

	Language *string `json:"language,omitempty"`
	// 开始时间的时区信息。时区信息，参考时区映射关系。

	TimeZoneID *string `json:"timeZoneID,omitempty"`
	// 录播类型。默认为禁用。 - 0: 禁用 - 1: 直播 - 2: 录播 - 3: 直播+录播

	RecordType *int32 `json:"recordType,omitempty"`
	// 主流直播地址，最大不超过255个字符。在录播类型为 :直播、直播+录播时有效。

	LiveAddress *string `json:"liveAddress,omitempty"`
	// 辅流直播地址，最大不超过255个字符。在录播类型为: 直播、直播+录播时有效。

	AuxAddress *string `json:"auxAddress,omitempty"`
	// 是否录制辅流，在录播类型为：录播、录播+直播时有效。  - 0: 不录制。  - 1: 录制。

	RecordAuxStream *int32 `json:"recordAuxStream,omitempty"`

	ConfConfigInfo *RestConfConfigDto `json:"confConfigInfo,omitempty"`
	// 录播鉴权方式，在录播类型为:录播、直播+录播时有效。 - 0: 可通过链接观看/下载。 - 1: 企业用户可观看/下载。 - 2: 与会者可观看/下载。

	RecordAuthType *int32 `json:"recordAuthType,omitempty"`
	// 是否使用云会议室召开预约会议。默认不使用云会议室。 - 0: 不使用云会议室。 - 1: 使用云会议室。

	VmrFlag *int32 `json:"vmrFlag,omitempty"`

	CycleParams *CycleParams `json:"cycleParams,omitempty"`
	// 用于识别用户开会时绑定的云会议室。最大长度不超过512个字符。 ID可以从云会议室管理->分页查询用户云会议室中获取id字段。 - 不为空，则用ID查询云会议室信息。 - 为空，则查用户所有云会议室，如果有个人云会议室，用个人云会议室ID；没有个人云会议室，取最小云会议室ID。

	VmrID *string `json:"vmrID,omitempty"`
	// 会议方数，会议最大与会人数限制

	ConcurrentParticipants *int32 `json:"concurrentParticipants,omitempty"`
}

func (o RestScheduleConfDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestScheduleConfDto struct{}"
	}

	return strings.Join([]string{"RestScheduleConfDto", string(data)}, " ")
}
