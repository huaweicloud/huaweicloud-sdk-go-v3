package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设备告警实体列表
type GetDeviceAlarmArrayObject struct {

	// 设备告警记录ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 设备告警名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 设备告警ID
	AlarmId *int32 `json:"alarm_id,omitempty" xml:"alarm_id"`

	// 设备告警等级，紧急告警(critical)，严重告警(major)，一般告警(minor)
	Level *string `json:"level,omitempty" xml:"level"`

	// 设备平台
	Platform *string `json:"platform,omitempty" xml:"platform"`

	// 设备告警的影响
	Impact *string `json:"impact,omitempty" xml:"impact"`

	// 设备告警详情内容
	Detail *string `json:"detail,omitempty" xml:"detail"`

	// 设备告警原因
	Reason *string `json:"reason,omitempty" xml:"reason"`

	// 设备告警处理建议
	DealSuggestion *string `json:"deal_suggestion,omitempty" xml:"deal_suggestion"`

	// 创建时间（时间戳）
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`
}

func (o GetDeviceAlarmArrayObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDeviceAlarmArrayObject struct{}"
	}

	return strings.Join([]string{"GetDeviceAlarmArrayObject", string(data)}, " ")
}
