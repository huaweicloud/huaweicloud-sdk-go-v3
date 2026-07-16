package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmConfigResponse Response Object
type ShowAlarmConfigResponse struct {
	Data *interface{} `json:"data,omitempty"`

	// 告警配置列表
	AlarmConfigs   *[]AlarmConfig `json:"alarm_configs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmConfigResponse", string(data)}, " ")
}
