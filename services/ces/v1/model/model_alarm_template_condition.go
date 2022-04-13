package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建自定义告警模板的告警策略。
type AlarmTemplateCondition struct {
	// 告警阈值的比较条件，可以是>、=、<、>=、<=。

	ComparisonOperator string `json:"comparison_operator"`
	// 触发告警的连续发生次数，取值范围[1, 5]。

	Count int32 `json:"count"`
	// 数据聚合的方式，支持max、min、average、sum、variance，分别表示最大值、最小值、平均值、求和值、方差值。

	Filter string `json:"filter"`
	// 告警条件判断周期，单位为秒，支持的值为1，300，1200，3600，14400，86400。说明：当period设置为1时，表示以原始的指标数据判断告警。当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为0。

	Period int32 `json:"period"`
	// 数据的单位，最大长度为32位。

	Unit *string `json:"unit,omitempty"`
	// 告警阈值，取值范围[0, Number.MAX_VALUE]，Number.MAX_VALUE值为1.7976931348623157e+108。具体阈值取值请参见附录中各服务监控指标中取值范围，如支持监控的服务列表中ECS的CPU使用率cpu_util取值范围可配置80。

	Value float64 `json:"value"`
	// 发送告警的周期，值可为0, 300, 600, 900, 1800, 3600, 10800, 21600, 43200, 86400；0表示只告警一次，300表示每5分钟告警一次，600表示每10分钟告警一次，900表示每15分钟告警一次，1800表示每30分钟告警一次，3600表示每1小时告警一次，10800表示每3小时告警一次，21600表示每6小时告警一次，43200表示每12小时告警一次，86400表示每1天告警一次。

	SuppressDuration *int32 `json:"suppress_duration,omitempty"`
}

func (o AlarmTemplateCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateCondition struct{}"
	}

	return strings.Join([]string{"AlarmTemplateCondition", string(data)}, " ")
}
