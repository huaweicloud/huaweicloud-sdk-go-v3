package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAlarmConfigsResponse struct {

	// 告警配置总数
	Count *int32 `json:"count,omitempty"`

	// 告警配置列表
	AlarmConfigs   *[]AlarmConfigResponse `json:"alarm_configs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAlarmConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmConfigsResponse", string(data)}, " ")
}
