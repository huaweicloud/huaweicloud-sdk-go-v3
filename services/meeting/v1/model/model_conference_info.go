package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 会议信息
type ConferenceInfo struct {

	// 会议ID。长度限制为32个字符。
	ConferenceID *string `json:"conferenceID,omitempty" xml:"conferenceID"`

	// 会议主题。长度限制为128个字符。
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 会议方数。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 时区参考。
	TimeZoneID *string `json:"timeZoneID,omitempty" xml:"timeZoneID"`

	// 会议起始时间 (YYYY-MM-DD HH:MM )。
	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	// 会议结束时间 (YYYY-MM-DD HH:MM )。
	EndTime *string `json:"endTime,omitempty" xml:"endTime"`

	// 会议的媒体类型。 由1个或多个枚举String组成，多个枚举时，每个枚举值之间通过”,”逗号分隔。 - Voice: 语音。 - Video: 标清视频。 - HDVideo: 高清视频（与Video互斥，如果同时选择Video、HDVideo，则系统默认选择Video）。 - Telepresence: 智真(与HDVideo、Video互斥，如果同时选择，系统使用Telepresence)。（预留字段） - Data: 多媒体。
	MediaTypes *string `json:"mediaTypes,omitempty" xml:"mediaTypes"`

	// 目前只会返回Created和Schedule状态， 如果会议已经召开返回Created状态，否则返回Schedule状态。 - Schedule: 预定状态。 - Creating: 正在创建状态。 - Created: 会议已经被创建，并正在召开。 - Destroyed: 会议已经关闭。
	ConferenceState *string `json:"conferenceState,omitempty" xml:"conferenceState"`

	// 会议语言。
	Language *string `json:"language,omitempty" xml:"language"`

	// 会议接入码。
	AccessNumber *string `json:"accessNumber,omitempty" xml:"accessNumber"`

	// 会议密码条目。预订者返回主持人密码和来宾密码。 - 主持人查询时返回主持人密码。 - 来宾查询时返回来宾密码。
	PasswordEntry *[]PasswordEntry `json:"passwordEntry,omitempty" xml:"passwordEntry"`

	// 会议预订者UUID。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 会议预订者帐号名称。长度最大限制为96个字符。
	ScheduserName *string `json:"scheduserName,omitempty" xml:"scheduserName"`

	// - 0: 普通会议。 - 2: 周期性会议。
	ConferenceType *int32 `json:"conferenceType,omitempty" xml:"conferenceType"`

	// 会议类型。 - FUTURE - IMMEDIATELY - CYCLE
	ConfType *string `json:"confType,omitempty" xml:"confType"`

	CycleParams *CycleParams `json:"cycleParams,omitempty" xml:"cycleParams"`

	// 是否入会自动静音。 - 0: 不自动静音 - 1: 自动静音
	IsAutoMute *int32 `json:"isAutoMute,omitempty" xml:"isAutoMute"`

	// 是否自动开启录音。 - 0: 不自动启动。 - 1: 自动启动。
	IsAutoRecord *int32 `json:"isAutoRecord,omitempty" xml:"isAutoRecord"`

	// 主持人会议链接地址。
	ChairJoinUri *string `json:"chairJoinUri,omitempty" xml:"chairJoinUri"`

	// 普通与会者会议链接地址。最大长度1024。
	GuestJoinUri *string `json:"guestJoinUri,omitempty" xml:"guestJoinUri"`

	// 旁听者会议链接地址。最大长度1024。（预留字段）
	AudienceJoinUri *string `json:"audienceJoinUri,omitempty" xml:"audienceJoinUri"`

	// 录播类型。 - 0: 禁用 。 - 1: 直播 。 - 2: 录播 。 - 3: 直播+录播。
	RecordType *int32 `json:"recordType,omitempty" xml:"recordType"`

	// 辅流直播地址。
	AuxAddress *string `json:"auxAddress,omitempty" xml:"auxAddress"`

	// 主流直播地址。
	LiveAddress *string `json:"liveAddress,omitempty" xml:"liveAddress"`

	// 是否录制辅流。  - 0: 否。  - 1: 是。
	RecordAuxStream *int32 `json:"recordAuxStream,omitempty" xml:"recordAuxStream"`

	// 录播鉴权方式。录播类型为:录播、直播+录播时有效。  - 0: 老的鉴权方式，url中携带token鉴权。  - 1: 企业内会议用户鉴权。  - 2: 会议内会议用户鉴权。
	RecordAuthType *int32 `json:"recordAuthType,omitempty" xml:"recordAuthType"`

	// 直播地址。（配置直播房间时会返回）
	LiveUrl *string `json:"liveUrl,omitempty" xml:"liveUrl"`

	ConfConfigInfo *RestConfConfigDto `json:"confConfigInfo,omitempty" xml:"confConfigInfo"`

	// 是否使用云会议室召开预约会议。 - 0: 不使用云会议室; - 1: 使用云会议室。 界面显示会议ID需要使用vmrConferenceID作为会议ID；查询会议详情、登录会控、一键入会等会议业务操作依然使用conferenceID字段。
	VmrFlag *int32 `json:"vmrFlag,omitempty" xml:"vmrFlag"`

	// 仅历史会议返回值有效。默认没有录制文件。 - True: 有录制文件。 - False: 没有录制文件。
	IsHasRecordFile *bool `json:"isHasRecordFile,omitempty" xml:"isHasRecordFile"`

	// 云会议室id，如果vmrFlag为1，则该字段不为空。
	VmrConferenceID *string `json:"vmrConferenceID,omitempty" xml:"vmrConferenceID"`

	// 会议的UUID。
	ConfUUID *string `json:"confUUID,omitempty" xml:"confUUID"`

	// 与会方信息。硬件终端/与会人最多各显示20条记录。
	PartAttendeeInfo *[]PartAttendee `json:"partAttendeeInfo,omitempty" xml:"partAttendeeInfo"`

	// 硬终端个数。
	TerminlCount *int32 `json:"terminlCount,omitempty" xml:"terminlCount"`

	// 普通终端个数。
	NormalCount *int32 `json:"normalCount,omitempty" xml:"normalCount"`

	// 会议预定者的企业名称。最大长度96。
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 云会议室的ID。
	VmrID *string `json:"vmrID,omitempty" xml:"vmrID"`

	// 会议角色
	Role *ConferenceInfoRole `json:"role,omitempty" xml:"role"`

	// 是否网络研讨会
	Webinar *bool `json:"webinar,omitempty" xml:"webinar"`

	// 标识是否为多流视频会议。 枚举值如下 1：多流会议
	MultiStreamFlag *int32 `json:"multiStreamFlag,omitempty" xml:"multiStreamFlag"`

	// 会议类型
	ConfMode *ConferenceInfoConfMode `json:"confMode,omitempty" xml:"confMode"`

	// True: VMR预约记录（如果为true则该记录不支持根据会议ID查询会议详情） False：普通会议
	ScheduleVmr *bool `json:"scheduleVmr,omitempty" xml:"scheduleVmr"`

	// 会议方数，会议最大与会人数限制
	ConcurrentParticipants *int32 `json:"concurrentParticipants,omitempty" xml:"concurrentParticipants"`

	PicDisplay *MultiPicDisplayDo `json:"picDisplay,omitempty" xml:"picDisplay"`

	// 周期子会议列表
	SubConfs *[]CycleSubConf `json:"subConfs,omitempty" xml:"subConfs"`

	// 周期子会议UUID, 用于查询在线会议和历史会议详情时标识
	CycleSubConfID *string `json:"cycleSubConfID,omitempty" xml:"cycleSubConfID"`
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
