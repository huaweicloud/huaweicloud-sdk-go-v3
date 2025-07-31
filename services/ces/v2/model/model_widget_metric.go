package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WidgetMetric struct {

	// 服务维度
	Namespace string `json:"namespace"`

	Dimensions *DimensionInfo `json:"dimensions"`

	// 多个指标名称，用逗号隔开
	MetricName string `json:"metric_name"`

	// 监控视图的指标别名列表
	Alias *[]string `json:"alias,omitempty"`

	ExtraInfo *ExtraInfo `json:"extra_info,omitempty"`

	// **参数解释** 是否开启聚合 **约束限制** 当RollupEnable开启时，RollupFilter和RollupDimension必填 **取值范围** true，表示开启聚合；false表示不开启聚合 **默认取值** false
	RollupEnable *bool `json:"rollup_enable,omitempty"`

	RollupFilter *RollupFilter `json:"rollup_filter,omitempty"`

	// 聚合维度
	RollupDimension *string `json:"rollup_dimension,omitempty"`

	// 是否展示同比（上周同一时间）数据，true:展示，false:不展示
	LastWeekCompareEnable *bool `json:"last_week_compare_enable,omitempty"`

	// 是否展示环比（昨天同一时间）数据，true:展示，false:不展示
	YesterdayCompareEnable *bool `json:"yesterday_compare_enable,omitempty"`

	// 维度名称，多维度用逗号分隔，各服务支持的维度可参考：“[服务维度名称](ces_03_0059.xml)”
	MetricDimension *string `json:"metric_dimension,omitempty"`

	// 展示数据数量
	TopNum *int32 `json:"top_num,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 排序字段，asc正序，desc倒序
	Order *WidgetMetricOrder `json:"order,omitempty"`

	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。
	TopnMetricName *string `json:"topn_metric_name,omitempty"`
}

func (o WidgetMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetMetric struct{}"
	}

	return strings.Join([]string{"WidgetMetric", string(data)}, " ")
}

type WidgetMetricOrder struct {
	value string
}

type WidgetMetricOrderEnum struct {
	ASC  WidgetMetricOrder
	DESC WidgetMetricOrder
}

func GetWidgetMetricOrderEnum() WidgetMetricOrderEnum {
	return WidgetMetricOrderEnum{
		ASC: WidgetMetricOrder{
			value: "asc",
		},
		DESC: WidgetMetricOrder{
			value: "desc",
		},
	}
}

func (c WidgetMetricOrder) Value() string {
	return c.value
}

func (c WidgetMetricOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WidgetMetricOrder) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
