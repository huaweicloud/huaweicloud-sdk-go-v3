package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议与会者QoS告警信息。
type QosParticipantInfo struct {

	// 会议的UUID。
	ConfUUID *string `json:"confUUID,omitempty"`

	// 会议ID。
	ConferenceID *string `json:"conferenceID,omitempty"`

	// 与会者标识。
	ParticipantID *string `json:"participantID,omitempty"`

	// 用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// 与会者的名称。
	DisplayName *string `json:"displayName,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`

	// 入会终端类型。 - PC: PC机 - MOBILE: 手机 - PAD：PAD设备 - MAC：MAC设备 - WEB：WEB方式入会，如通过WebRTC入会 - ROOM: 会议室 - 硬件终端：显示具体的硬件设备类型，如TE50, HUAWEI IDEAHUB, CISCO等 - OTHER: 其他设备
	TerminalType *string `json:"terminalType,omitempty"`

	// 与会者角色。 - host：主持人 - guest：来宾 - audience：观众
	Role *string `json:"role,omitempty"`

	// 与会者的IP地址。
	IpAddress *string `json:"ipAddress,omitempty"`

	// 国家。
	Country *string `json:"country,omitempty"`

	// 省市（仅限中国）。
	Province *string `json:"province,omitempty"`

	// 城市（仅限中国）。
	City *string `json:"city,omitempty"`

	// 华为云会议APP版本。
	AppVersion *string `json:"appVersion,omitempty"`

	// 入会时间(UTC时间), Unix时间戳（单位毫秒）。
	JoinTime *int64 `json:"joinTime,omitempty"`

	// 离会时间(UTC时间), Unix时间戳（单位毫秒）。 > * 与会者未离会：leftTime = 0 > * 与会者已离会：leftTime = 实际离会时间
	LeftTime *int64 `json:"leftTime,omitempty"`

	// 终端操作系统信息。
	SystemInfo *string `json:"systemInfo,omitempty"`

	// 网络类型。
	NetworkType *string `json:"networkType,omitempty"`

	// 总体告警。 * YES：音频（发送/接收），视频（发送/接收），屏幕共享（发送/接收），CPU任一项产生告警，总体告警状态即为 YES * NO：无告警
	Alarm *string `json:"alarm,omitempty"`

	// 音频发送告警。 * YES ：发送音频的抖动，时延，丢包率任一项产生阈值告警，则音频发送告警状态为YES * NO：发送音频无告警
	AudioAlarmSend *string `json:"audioAlarmSend,omitempty"`

	// 视频发送告警。 * YES ：发送视频的抖动，时延，丢包率任一项产生阈值告警，则视频发送告警状态为YES * NO：发送视频无告警
	VideoAlarmSend *string `json:"videoAlarmSend,omitempty"`

	// 屏幕共享发送告警。 * YES：发送屏幕共享的抖动，时延，丢包率任一项产生阈值告警，则屏幕共享发送告警状态为YES * NO：发送屏幕共享无告警
	ScreenAlarmSend *string `json:"screenAlarmSend,omitempty"`

	// 音频接收告警。 * YES：接收音频的抖动，时延，丢包率任一项产生阈值告警，则音频接收告警状态为YES * NO：接收音频无告警
	AudioAlarmRec *string `json:"audioAlarmRec,omitempty"`

	// 视频接收告警。 * YES：接收视频的抖动，时延，丢包率任一项产生阈值告警，则视频接收告警状态为YES * NO：接收视频无告警
	VideoAlarmRec *string `json:"videoAlarmRec,omitempty"`

	// 屏幕共享接收告警。 * YES：接收屏幕共享的抖动，时延，丢包率任一项产生阈值告警，则屏幕共享接收告警状态为YES * NO：接收屏幕共享无告警
	ScreenAlarmRec *string `json:"screenAlarmRec,omitempty"`

	// CPU告警。 * YES：端侧的APP最大CPU使用率或系统最大CPU使用率任一项产生阈值告警，则CPU告警状态为YES * NO：CPU无告警
	CpuAlarm *string `json:"cpuAlarm,omitempty"`

	// 麦克风。
	MicrophoneInfo *string `json:"microphoneInfo,omitempty"`

	// 扬声器。
	SpeakerInfo *string `json:"speakerInfo,omitempty"`

	// 摄像头。
	CameraInfo *string `json:"cameraInfo,omitempty"`

	// 数据中心。
	DataCenter *string `json:"dataCenter,omitempty"`

	// 离会原因。此字段仅标识离会原因，不做为是否已离会的判断依据。正在与会人员的离会原因初始值 = 0。 * 0：正常离会 * 1：网络异常离会
	LeftReason *int32 `json:"leftReason,omitempty"`

	// 与会者是否存在QoS数据。 * true：存在QoS数据 * false：不存在QoS数据
	ExistQos *bool `json:"existQos,omitempty"`
}

func (o QosParticipantInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosParticipantInfo struct{}"
	}

	return strings.Join([]string{"QosParticipantInfo", string(data)}, " ")
}
