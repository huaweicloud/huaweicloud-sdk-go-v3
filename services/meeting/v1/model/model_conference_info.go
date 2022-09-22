package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 会议信息。
type ConferenceInfo struct {

	// 会议ID。
	ConferenceID *string `json:"conferenceID,omitempty"`

	// 会议主题。
	Subject *string `json:"subject,omitempty"`

	// 会议预约时添加的会议者数量。
	Size *int32 `json:"size,omitempty"`

	// 会议通知中会议时间的时区信息。时区信息，参考[[时区映射关系](https://support.huaweicloud.com/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hws)[[时区映射关系](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0110.html#ZH-CN_TOPIC_0212714472__table137407441463)](tag:hk)。
	TimeZoneID *string `json:"timeZoneID,omitempty"`

	// 会议起始时间 (YYYY-MM-DD HH:MM )。
	StartTime *string `json:"startTime,omitempty"`

	// 会议结束时间 (YYYY-MM-DD HH:MM )。
	EndTime *string `json:"endTime,omitempty"`

	// 会议的媒体类型。 由1个或多个枚举String组成，多个枚举时，每个枚举值之间通过”,”逗号分隔。 - Voice: 语音 - Video: 标清视频 - HDVideo: 高清视频 - Data: 数据
	MediaTypes *string `json:"mediaTypes,omitempty"`

	// 会议状态。 - Schedule: 预定状态 - Creating: 正在创建状态 - Created: 会议已经被创建，并正在召开 - Destroyed: 会议已经关闭
	ConferenceState *string `json:"conferenceState,omitempty"`

	// 会议通知短信或邮件的语言。默认中文。 * zh-CN：中文 * en-US：英文
	Language *string `json:"language,omitempty"`

	// 会议接入的SIP号码。
	AccessNumber *string `json:"accessNumber,omitempty"`

	// 会议密码。 > * 创建会议时，返回主持人密码和来宾密码 > * 主持人查询会议时，返回主持人密码和来宾密码来 > * 宾查询会议时，返回来宾密码
	PasswordEntry *[]PasswordEntry `json:"passwordEntry,omitempty"`

	// 会议预订者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// 会议预订者名称。
	ScheduserName *string `json:"scheduserName,omitempty"`

	// 会议类型。 - 0: 普通会议 - 2: 周期性会议
	ConferenceType *int32 `json:"conferenceType,omitempty"`

	// 会议类型。 - FUTURE：将来开始的会议（创建时） - IMMEDIATELY：立即开始的会议（创建时） - CYCLE：周期会议
	ConfType *string `json:"confType,omitempty"`

	CycleParams *CycleParams `json:"cycleParams,omitempty"`

	// 是否入会自动静音。 - 0: 不自动静音 - 1: 自动静音
	IsAutoMute *int32 `json:"isAutoMute,omitempty"`

	// 是否自动开启云录制。 - 0: 不自动启动 - 1: 自动启动
	IsAutoRecord *int32 `json:"isAutoRecord,omitempty"`

	// 主持人会议链接地址。
	ChairJoinUri *string `json:"chairJoinUri,omitempty"`

	// 普通与会者会议链接地址。
	GuestJoinUri *string `json:"guestJoinUri,omitempty"`

	// 网络研讨会观众会议链接地址。
	AudienceJoinUri *string `json:"audienceJoinUri,omitempty"`

	// 录播类型。 - 0: 禁用 - 1: 直播 - 2: 录播 - 3: 直播+录播
	RecordType *int32 `json:"recordType,omitempty"`

	// 辅流直播推流地址。
	AuxAddress *string `json:"auxAddress,omitempty"`

	// 主流直播推流地址。
	LiveAddress *string `json:"liveAddress,omitempty"`

	// 是否录制辅流。  - 0: 否  - 1: 是
	RecordAuxStream *int32 `json:"recordAuxStream,omitempty"`

	// 观看/下载录播的鉴权方式。  - 0: 可通过链接观看/下载  - 1: 企业用户可观看/下载  - 2: 与会者可观看/下载
	RecordAuthType *int32 `json:"recordAuthType,omitempty"`

	// 直播观看地址。
	LiveUrl *string `json:"liveUrl,omitempty"`

	ConfConfigInfo *RestConfConfigDto `json:"confConfigInfo,omitempty"`

	// 是否使用云会议室或个人会议ID召开预约会议。 - 0: 不使用云会议室或个人会议ID - 1: 使用云会议室或个人会议ID
	VmrFlag *int32 `json:"vmrFlag,omitempty"`

	// 是否有会议录制文件。仅历史会议查询时返回。 - true: 有录制文件 - false: 没有录制文件
	IsHasRecordFile *bool `json:"isHasRecordFile,omitempty"`

	// 云会议室会议ID或个人会议ID，如果vmrFlag为\"1\"，则该字段不为空。
	VmrConferenceID *string `json:"vmrConferenceID,omitempty"`

	// 会议的UUID。 > * 只有创建立即开始的会议才返回UUID，如果是预约未来的会议，不会返回UUID > * 可以通过[[查询历史会议列表](https://support.huaweicloud.com/api-meeting/meeting_21_0051.html)](tag:hws)[[查询历史会议列表](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0051.html)](tag:hk)获取历史会议的UUID
	ConfUUID *string `json:"confUUID,omitempty"`

	// 被邀请的部分与会者信息。 > * 只返回被邀请的前20条软终端与会者信息和前20条硬终端与会者信息 > * 不返回会中主动加入的与会者信息 > * “[[查询会议列表](https://support.huaweicloud.com/api-meeting/meeting_21_0017.html)](tag:hws)[[查询会议列表](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0017.html)](tag:hk)”和“[[查询会议详情](https://support.huaweicloud.com/api-meeting/meeting_21_0018.html)](tag:hws)[[查询会议详情](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0018.html)](tag:hk)”接口，返回预约会议时邀请的与会者和会中主持人邀请的与会者 > * “[[查询在线会议列表](https://support.huaweicloud.com/api-meeting/meeting_21_0025.html)](tag:hws)[[查询在线会议列表](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0025.html)](tag:hk)”、“[[查询在线会议详情](https://support.huaweicloud.com/api-meeting/meeting_21_0026.html)](tag:hws)[[查询在线会议详情](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0026.html)](tag:hk)”、“[[查询历史会议列表](https://support.huaweicloud.com/api-meeting/meeting_21_0051.html)](tag:hws)[[查询历史会议列表](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0051.html)](tag:hk)”和“[[查询历史会议详情](https://support.huaweicloud.com/api-meeting/meeting_21_0052.html)](tag:hws)[[查询历史会议详情](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0052.html)](tag:hk)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       ”接口返回预约会议时邀请的与会者。不返回会中主持人邀请的与会者
	PartAttendeeInfo *[]PartAttendee `json:"partAttendeeInfo,omitempty"`

	// 硬终端个数，如IdeaHub，TE30等。
	TerminlCount *int32 `json:"terminlCount,omitempty"`

	// 软终端个数，如PC端、手机端App等。
	NormalCount *int32 `json:"normalCount,omitempty"`

	// 会议预定者的企业名称。
	DeptName *string `json:"deptName,omitempty"`

	// 云会议室的ID。
	VmrID *string `json:"vmrID,omitempty"`

	// 与会者角色。 * chair ：主持人 * general ：来宾 * audience ： 观众 > * 仅在查询会议详情时返回 > * 返回查询者本身的角色
	Role *ConferenceInfoRole `json:"role,omitempty"`

	// 是否是网络研讨会。
	Webinar *bool `json:"webinar,omitempty"`

	// 标识是否为多流视频会议。 * 1：多流会议
	MultiStreamFlag *int32 `json:"multiStreamFlag,omitempty"`

	// 会议类型模型。 * COMMON：MCU会议 * RTC：MMR会议
	ConfMode *ConferenceInfoConfMode `json:"confMode,omitempty"`

	// VMR预约记录。 true: VMR预约记录 false：普通会议 > 该参数将废弃，请勿使用。
	ScheduleVmr *bool `json:"scheduleVmr,omitempty"`

	// 会议最大与会人数。默认值0。 * 0：无限制 * 大于0：会议最大与会人数
	ConcurrentParticipants *int32 `json:"concurrentParticipants,omitempty"`

	PicDisplay *MultiPicDisplayDo `json:"picDisplay,omitempty"`

	// 周期子会议列表。
	SubConfs *[]CycleSubConf `json:"subConfs,omitempty"`

	// 第一个周期子会议的UUID。
	CycleSubConfID *string `json:"cycleSubConfID,omitempty"`
}

func (o ConferenceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConferenceInfo struct{}"
	}

	return strings.Join([]string{"ConferenceInfo", string(data)}, " ")
}

type ConferenceInfoRole struct {
	value string
}

type ConferenceInfoRoleEnum struct {
	CHAIR    ConferenceInfoRole
	GENERAL  ConferenceInfoRole
	AUDIENCE ConferenceInfoRole
}

func GetConferenceInfoRoleEnum() ConferenceInfoRoleEnum {
	return ConferenceInfoRoleEnum{
		CHAIR: ConferenceInfoRole{
			value: "chair",
		},
		GENERAL: ConferenceInfoRole{
			value: "general",
		},
		AUDIENCE: ConferenceInfoRole{
			value: "audience",
		},
	}
}

func (c ConferenceInfoRole) Value() string {
	return c.value
}

func (c ConferenceInfoRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConferenceInfoRole) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ConferenceInfoConfMode struct {
	value string
}

type ConferenceInfoConfModeEnum struct {
	COMMON ConferenceInfoConfMode
	RTC    ConferenceInfoConfMode
}

func GetConferenceInfoConfModeEnum() ConferenceInfoConfModeEnum {
	return ConferenceInfoConfModeEnum{
		COMMON: ConferenceInfoConfMode{
			value: "COMMON",
		},
		RTC: ConferenceInfoConfMode{
			value: "RTC",
		},
	}
}

func (c ConferenceInfoConfMode) Value() string {
	return c.value
}

func (c ConferenceInfoConfMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConferenceInfoConfMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
