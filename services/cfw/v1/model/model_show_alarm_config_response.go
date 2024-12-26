package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmConfigResponse Response Object
type ShowAlarmConfigResponse struct {

	// 告警配置列表
	AlarmConfigs *[]AlarmConfig `json:"alarm_configs,omitempty"`

	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmConfigResponse", string(data)}, " ")
}
