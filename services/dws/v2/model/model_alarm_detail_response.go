package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警详情
type AlarmDetailResponse struct {

	// 告警定义ID
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警名称
	AlarmName *string `json:"alarm_name,omitempty"`

	// 告警级别
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// 告警服务
	AlarmSource *string `json:"alarm_source,omitempty"`

	// 告警消息
	AlarmMessage *string `json:"alarm_message,omitempty"`

	// 告警定位信息
	AlarmLocation *string `json:"alarm_location,omitempty"`

	// 告警源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 告警源名称
	ResourceIdName *string `json:"resource_id_name,omitempty"`

	// 告警日期
	AlarmGenerateDate *string `json:"alarm_generate_date,omitempty"`

	// 告警状态
	AlarmStatus *string `json:"alarm_status,omitempty"`
}

func (o AlarmDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDetailResponse struct{}"
	}

	return strings.Join([]string{"AlarmDetailResponse", string(data)}, " ")
}
