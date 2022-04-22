package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警触发条件
type AlarmCondition struct {

	// 告警条件判断周期，单位为秒，支持的值为1，300，1200，3600，14400，86400。说明：当period设置为1时，表示以原始的指标数据判断告警。当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为0。
	Period int32 `json:"period"`

	// 聚合方式
	Filter string `json:"filter"`

	// 阈值符号
	ComparisonOperator string `json:"comparison_operator"`

	// 告警阈值，取值范围[0, Number.MAX_VALUE]，Number.MAX_VALUE值为1.7976931348623157e+108。具体阈值取值请参见附录中各服务监控指标中取值范围，如支持监控的服务列表中ECS的CPU使用率cpu_util取值范围可配置80。
	Value float32 `json:"value"`

	// 数据的单位，最大长度为32位。
	Unit *string `json:"unit,omitempty"`

	// 次数
	Count int32 `json:"count"`

	// 发送告警的周期，值可为0, 300, 600, 900, 1800, 3600, 10800, 21600, 43200, 86400；0表示只告警一次，300表示每5分钟告警一次，600表示每10分钟告警一次，900表示每15分钟告警一次，1800表示每30分钟告警一次，3600表示每1小时告警一次，10800表示每3小时告警一次，21600表示每6小时告警一次，43200表示每12小时告警一次，86400表示每1天告警一次。
	SuppressDuration *int32 `json:"suppress_duration,omitempty"`
}

func (o AlarmCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmCondition struct{}"
	}

	return strings.Join([]string{"AlarmCondition", string(data)}, " ")
}
