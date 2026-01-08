package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WidgetMetric struct {

	// **参数解释** 查询服务的命名空间，各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg) **约束限制** 不涉及 **取值范围** 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值** 不涉及
	Namespace string `json:"namespace"`

	Dimensions *DimensionInfo `json:"dimensions"`

	// **参数解释** 多个指标名称 **约束限制** 不涉及 **取值范围** 长度为[1,1080]个字符，多个指标名称之间用逗号隔开 **默认取值** 不涉及
	MetricName string `json:"metric_name"`

	// **参数解释** 监控视图的指标别名列表 **约束限制** 当资源类型为指定资源时才允许传该参数 包含的指标别名对象个数为[0,200]
	Alias *[]string `json:"alias,omitempty"`

	ExtraInfo *ExtraInfo `json:"extra_info,omitempty"`

	// **参数解释** 是否开启聚合 **约束限制** 当rollup_enable开启时，rollup_filter和rollup_dimension必填 **取值范围** - true：表示开启聚合 - false：表示不开启聚合 **默认取值** false
	RollupEnable *bool `json:"rollup_enable,omitempty"`

	RollupFilter *RollupFilter `json:"rollup_filter,omitempty"`

	// **参数解释** 聚合维度 **约束限制** 不涉及 **取值范围** 长度为[1,32]个字符 **默认取值** 不涉及
	RollupDimension *string `json:"rollup_dimension,omitempty"`

	// **参数解释** 是否展示同比（上周同一时间）数据 **约束限制** 不涉及 **取值范围** - true:展示 - false:不展示 **默认取值** 不涉及
	LastWeekCompareEnable *bool `json:"last_week_compare_enable,omitempty"`

	// **参数解释** 是否展示环比（昨天同一时间）数据 **约束限制** 不涉及 **取值范围** - true:展示 - false:不展示 **默认取值** 不涉及
	YesterdayCompareEnable *bool `json:"yesterday_compare_enable,omitempty"`

	// **参数解释** 维度名称，多维度用逗号分隔，各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。，必须以字母开头，只能包含0-9/a-z/A-Z/_/-，多维度用\",\"分隔，每个维度的最大长度为32。总长度为[1,131]个字符。目前最大支持4个维度。举例：单维度场景：instance_id；多维度场景：instance_id,disk **约束限制** 不涉及           **取值范围** 长度为[1,131]个字符 **默认取值** 不涉及
	MetricDimension *string `json:"metric_dimension,omitempty"`

	// **参数解释** 展示数据数量 **约束限制** 不涉及                 **取值范围** 最小值为1，最大值为200 **默认取值** 不涉及
	TopNum *int32 `json:"top_num,omitempty"`

	// **参数解释** 单位 **约束限制** 不涉及 **取值范围** 长度为[0,32]个字符 **默认取值** 不涉及
	Unit *string `json:"unit,omitempty"`

	// **参数解释** 排序字段 **约束限制** 不涉及                **取值范围** - asc:正序 - desc:倒序 **默认取值** 不涉及
	Order *WidgetMetricOrder `json:"order,omitempty"`

	// **参数解释** 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务资源的指标名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **约束限制** 不涉及 **取值范围** 长度为[1,96]个字符 **默认取值** 不涉及
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
