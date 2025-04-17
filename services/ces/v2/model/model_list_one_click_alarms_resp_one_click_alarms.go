package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListOneClickAlarmsRespOneClickAlarms struct {

	// 一键告警ID
	OneClickAlarmId string `json:"one_click_alarm_id"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace string `json:"namespace"`

	// 一键告警描述，长度范围[0,256]，该字段默认值为空字符串
	Description string `json:"description"`

	// 开关
	Enabled bool `json:"enabled"`
}

func (o ListOneClickAlarmsRespOneClickAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOneClickAlarmsRespOneClickAlarms struct{}"
	}

	return strings.Join([]string{"ListOneClickAlarmsRespOneClickAlarms", string(data)}, " ")
}
