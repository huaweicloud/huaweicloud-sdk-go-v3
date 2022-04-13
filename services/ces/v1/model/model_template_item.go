package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建自定义告警模板添加的告警规则。
type TemplateItem struct {
	// 告警模板添加的监控指标，如弹性云服务器可添加的监控指标为cpu_util等；各服务的指标名称可查看：“[服务指标名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	MetricName string `json:"metric_name"`

	Condition *AlarmTemplateCondition `json:"condition"`
	// 设置告警级别，值为1,2,3,4；1为紧急，2为重要，3为次要，4为提示。

	AlarmLevel *int32 `json:"alarm_level,omitempty"`
}

func (o TemplateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateItem struct{}"
	}

	return strings.Join([]string{"TemplateItem", string(data)}, " ")
}
