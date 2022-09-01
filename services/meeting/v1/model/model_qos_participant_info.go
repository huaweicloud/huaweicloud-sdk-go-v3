package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 与会者信息。
type QosParticipantInfo struct {

	// 会议的UUID。
	ConfUUID *string `json:"confUUID,omitempty" xml:"confUUID"`

	// 会议ID。
	ConferenceID *string `json:"conferenceID,omitempty" xml:"conferenceID"`

	// 入会UUID。
	ParticipantID *string `json:"participantID,omitempty" xml:"participantID"`

	// 用户UUID。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 与会者的名称（昵称）。
	DisplayName *string `json:"displayName,omitempty" xml:"displayName"`

	// 部门。
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 入会终端类型。 - PC: PC机。 - MOBILE: 手机。 - PAD：PAD设备。 - MAC：MAC设备。 - WEB：WEB方式入会，如通过WebRTC入会。 - ROOM: 会议室。 - 硬件终端：显示具体的硬件设备类型，如TE50, HUAWEI IDEAHUB, CISCO等。 - OTHER: 其他设备。
	TerminalType *string `json:"terminalType,omitempty" xml:"terminalType"`

	// 与会者角色。 - host：主持人。 - guest：来宾。 - audience：观众。
	Role *string `json:"role,omitempty" xml:"role"`

	// 与会者的IP地址。
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress"`

	// 国家。
	Country *string `json:"country,omitempty" xml:"country"`

	// 省市（仅限中国）。
	Province *string `json:"province,omitempty" xml:"province"`

	// 城市（仅限中国）。
	City *string `json:"city,omitempty" xml:"city"`

	// 华为云会议APP版本。
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion"`

	// 入会时间(UTC时间), Unix时间戳（单位毫秒）。
	JoinTime *int64 `json:"joinTime,omitempty" xml:"joinTime"`

	// 离会时间(UTC时间), Unix时间戳（单位毫秒）。 说明： * 与会者未离会：leftTime = 0。 * 与会者已离会：leftTime = 实际离会时间。
	LeftTime *int64 `json:"leftTime,omitempty" xml:"leftTime"`

	// 系统信息。
	SystemInfo *string `json:"systemInfo,omitempty" xml:"systemInfo"`

	// 网络类型。
	NetworkType *string `json:"networkType,omitempty" xml:"networkType"`

	// 总体告警 YES/NO。 说明： * 音频（发送/接收），视频（发送/接收），屏幕共享（发送/接收），CPU任一项产生告警，总体告警状态即为 YES。
	Alarm *string `json:"alarm,omitempty" xml:"alarm"`

	// 音频发送告警 YES / NO。 说明： * 发送音频的抖动，时延，丢包率任一项产生阈值告警，则音频发送告警状态为YES。
	AudioAlarmSend *string `json:"audioAlarmSend,omitempty" xml:"audioAlarmSend"`

	// 视频发送告警 YES / NO。 说明： * 发送视频的抖动，时延，丢包率，分辨率任一项产生阈值告警，则视频发送告警状态为YES。
	VideoAlarmSend *string `json:"videoAlarmSend,omitempty" xml:"videoAlarmSend"`

	// 屏幕共享发送告警 YES / NO。 说明： * 发送屏幕共享的抖动，时延，丢包率任一项产生阈值告警，则屏幕共享发送告警状态为YES。
	ScreenAlarmSend *string `json:"screenAlarmSend,omitempty" xml:"screenAlarmSend"`

	// 音频接收告警 YES / NO。 说明： * 接收音频的抖动，时延，丢包率任一项产生阈值告警，则音频接收告警状态为YES。
	AudioAlarmRec *string `json:"audioAlarmRec,omitempty" xml:"audioAlarmRec"`

	// 视频接收告警 YES / NO。 说明： * 接收视频的抖动，时延，丢包率任一项产生阈值告警，则视频接收告警状态为YES。
	VideoAlarmRec *string `json:"videoAlarmRec,omitempty" xml:"videoAlarmRec"`

	// 屏幕共享接收告警 YES / NO。 说明： * 接收屏幕共享的抖动，时延，丢包率任一项产生阈值告警，则屏幕共享接收告警状态为YES。
	ScreenAlarmRec *string `json:"screenAlarmRec,omitempty" xml:"screenAlarmRec"`

	// CPU告警 YES / NO。 说明： * 端侧的APP最大CPU使用率或系统最大CPU使用率任一项产生阈值告警，则CPU告警状态为YES。
	CpuAlarm *string `json:"cpuAlarm,omitempty" xml:"cpuAlarm"`

	// 麦克风。
	MicrophoneInfo *string `json:"microphoneInfo,omitempty" xml:"microphoneInfo"`

	// 扬声器。
	SpeakerInfo *string `json:"speakerInfo,omitempty" xml:"speakerInfo"`

	// 摄像头。
	CameraInfo *string `json:"cameraInfo,omitempty" xml:"cameraInfo"`

	// 数据中心。
	DataCenter *string `json:"dataCenter,omitempty" xml:"dataCenter"`

	// 离会原因。此字段仅标识离会原因，不做为是否已离会的判断依据。正在与会人员的离会原因初始值 = 0。 说明： * 0：正常离会。 * 1：网络异常离会。
	LeftReason *int32 `json:"leftReason,omitempty" xml:"leftReason"`

	// 与会者是否存在QoS数据。 true：存在QoS数据。 false：不存在QoS数据。
	ExistQos *bool `json:"existQos,omitempty" xml:"existQos"`
}

func (o QosParticipantInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosParticipantInfo struct{}"
	}

	return strings.Join([]string{"QosParticipantInfo", string(data)}, " ")
}
