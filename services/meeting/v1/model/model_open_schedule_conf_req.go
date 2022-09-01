package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 网路研讨会议预定请求。
type OpenScheduleConfReq struct {

	// 主题
	Subject string `json:"subject" xml:"subject"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 会议开始时间（UTC时间）， 格式：yyyy-MM-dd HH:mm。
	StartTime string `json:"startTime" xml:"startTime"`

	// 会议持续时长，单位分钟，取值范围[15,1440]。
	Duration int32 `json:"duration" xml:"duration"`

	// 开始时间的时区信息。时区信息，参考时区映射关系。
	TimeZoneId int32 `json:"timeZoneId" xml:"timeZoneId"`

	// 与会者列表。该列表可以用于发送会议通知、会议提醒、会议开始时候进行自动邀请。
	Attendees *[]OpenAttendeeEntity `json:"attendees,omitempty" xml:"attendees"`

	NotifySetting *OpenNotifySetting `json:"notifySetting,omitempty" xml:"notifySetting"`

	// VMR ID, 用于识别用户开会时绑定的云会议室。最大长度不超过128个字符。
	VmrID string `json:"vmrID" xml:"vmrID"`

	// 自定义嘉宾入会密码, 4-16位数字，不能与观众密码相同；不指定则系统自动创建。
	GuestPasswd *string `json:"guestPasswd,omitempty" xml:"guestPasswd"`

	// 自定义观众入会密码, 4-16位数字，不能与嘉宾密码相同；不指定则系统自动创建。
	AudiencePasswd *string `json:"audiencePasswd,omitempty" xml:"audiencePasswd"`

	// 入会范围开关
	CallRestriction *bool `json:"callRestriction,omitempty" xml:"callRestriction"`

	// 主持人、嘉宾入会范围  0: 所有用户 1: 非匿名用户（手机pstn入会视为匿名入会） 2: 企业内用户 3: 被邀请用户; 默认值 0。
	Scope *int32 `json:"scope,omitempty" xml:"scope"`

	// 观众入会范围 0: 所有用户 2: 企业内用户和被邀请用户; 默认值 0。
	AudienceScope *int32 `json:"audienceScope,omitempty" xml:"audienceScope"`
}

func (o OpenScheduleConfReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenScheduleConfReq struct{}"
	}

	return strings.Join([]string{"OpenScheduleConfReq", string(data)}, " ")
}
