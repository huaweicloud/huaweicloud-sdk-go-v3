package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WidgetMetricResp struct {

	// **参数解释** 查询服务的命名空间，各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg) **取值范围** 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。
	Namespace *string `json:"namespace,omitempty"`

	Dimensions *DimensionInfoResp `json:"dimensions,omitempty"`

	// **参数解释** 多个指标名称 **取值范围** 长度为[1,1080]个字符，多个指标名称之间用逗号隔开
	MetricName *string `json:"metric_name,omitempty"`

	// **参数解释** 监控视图的指标别名列表
	Alias *[]string `json:"alias,omitempty"`

	ExtraInfo *ExtraInfoResp `json:"extra_info,omitempty"`

	// **参数解释** 是否开启聚合 **取值范围** - true：表示开启聚合 - false：表示不开启聚合
	RollupEnable *bool `json:"rollup_enable,omitempty"`

	RollupFilter *RollupFilterResp `json:"rollup_filter,omitempty"`

	// **参数解释** 聚合维度 **取值范围** 长度为[1,32]个字符
	RollupDimension *string `json:"rollup_dimension,omitempty"`

	// **参数解释** 是否展示同比（上周同一时间）数据 **取值范围** - true:展示 - false:不展示
	LastWeekCompareEnable *bool `json:"last_week_compare_enable,omitempty"`

	// **参数解释** 是否展示环比（昨天同一时间）数据 **取值范围** - true:展示 - false:不展示
	YesterdayCompareEnable *bool `json:"yesterday_compare_enable,omitempty"`

	// **参数解释** 维度名称，多维度用逗号分隔，各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。          **取值范围** 必须以字母开头，只能包含0-9/a-z/A-Z/_/-，多维度用\",\"分隔，每个维度的最大长度为32。总长度为[1,131]个字符。目前最大支持4个维度。举例：单维度场景：instance_id；多维度场景：instance_id,disk
	MetricDimension *string `json:"metric_dimension,omitempty"`

	// **参数解释** 展示数据数量               **取值范围** 最小值为1，最大值为200
	TopNum *int32 `json:"top_num,omitempty"`

	// **参数解释** 单位 **取值范围** 长度为[0,32]个字符
	Unit *string `json:"unit,omitempty"`

	// **参数解释** 排序字段               **取值范围** - asc:正序 - desc:倒序
	Order *WidgetMetricRespOrder `json:"order,omitempty"`

	// **参数解释** 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务资源的指标名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **取值范围** 长度为[1,96]个字符
	TopnMetricName *string `json:"topn_metric_name,omitempty"`
}

func (o WidgetMetricResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetMetricResp struct{}"
	}

	return strings.Join([]string{"WidgetMetricResp", string(data)}, " ")
}

type WidgetMetricRespOrder struct {
	value string
}

type WidgetMetricRespOrderEnum struct {
	ASC  WidgetMetricRespOrder
	DESC WidgetMetricRespOrder
}

func GetWidgetMetricRespOrderEnum() WidgetMetricRespOrderEnum {
	return WidgetMetricRespOrderEnum{
		ASC: WidgetMetricRespOrder{
			value: "asc",
		},
		DESC: WidgetMetricRespOrder{
			value: "desc",
		},
	}
}

func (c WidgetMetricRespOrder) Value() string {
	return c.value
}

func (c WidgetMetricRespOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WidgetMetricRespOrder) UnmarshalJSON(b []byte) error {
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
