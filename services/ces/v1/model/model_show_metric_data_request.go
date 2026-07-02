package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMetricDataRequest Request Object
type ShowMetricDataRequest struct {

	// **参数解释** 服务命名空间，样例：弹性云服务器的命名空间为SYS.ECS。 各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/zh-cn/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg) **约束限制** 不涉及 **取值范围** 格式为service.item，service和item以点号拼接组成。 service和item必须以字母（大写或小写）开头，后面可以跟零个或多个字母（大写或小写）、数字、下划线（_），长度为[3,32]个字符 **默认取值** 不涉及
	Namespace string `json:"namespace"`

	// **参数解释** 资源的监控指标名称，样例：弹性云服务器监控指标中的cpu_util。 各服务资源的指标名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/zh-cn/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档 **约束限制** 不涉及 **取值范围** 必须以字母（大写或小写）开头，后面可以跟零个或多个字母（大写或小写）、数字、下划线（_）、连字符 （-），长度为[1,96]个字符 **默认取值** 不涉及
	MetricName string `json:"metric_name"`

	// **参数解释** 指标的第一层维度，目前最多支持4个维度，维度格式为dim.{i}=key,value。样例：instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d 各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/zh-cn/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **约束限制** 不涉及 **取值范围** dim.0=key,value，由key、value以逗号拼接组成。 key必须以字母（大写或小写）开头，后面可以跟零个或多个字母（大写或小写）、数字、下划线（_）、连字符（-），长度为[1,32]个字符 value由多个字母（大写或小写）、数字、下划线（_）、连字符（-）、点（.）、斜杠（/）、井号（#）、英文左括号（(）、英文右括号（)）组合而成，首个字符可以包含星号（*），但不能以连字符（-）开头，长度为[1,256]个字符 **默认取值** 不涉及
	Dim0 string `json:"dim.0"`

	// **参数解释** 指标的第二层维度，目前最多支持4个维度，维度格式为dim.{i}=key,value。样例：instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d 各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/zh-cn/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **约束限制** 不涉及 **取值范围** dim.1=key,value，由key、value以逗号拼接组成。 key必须以字母（大写或小写）开头，后面可以跟零个或多个字母（大写或小写）、数字、下划线（_）、连字符（-），长度为[1,32]个字符 value由多个字母（大写或小写）、数字、下划线（_）、连字符（-）、点（.）、斜杠（/）、井号（#）、英文左括号（(）、英文右括号（)）组合而成，首个字符可以包含星号（*），但不能以连字符（-）开头，长度为[1,256]个字符 **默认取值** 不涉及
	Dim1 *string `json:"dim.1,omitempty"`

	// **参数解释** 指标的第三层维度，目前最多支持4个维度，维度格式为dim.{i}=key,value。样例：instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d 各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/zh-cn/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档 **约束限制** 不涉及 **取值范围** dim.2=key,value，由key、value以逗号拼接组成。 key必须以字母（大写或小写）开头，后面可以跟零个或多个字母（大写或小写）、数字、下划线（_）、连字符（-），长度为[1,32]个字符 value由多个字母（大写或小写）、数字、下划线（_）、连字符（-）、点（.）、斜杠（/）、井号（#）、英文左括号（(）、英文右括号（)）组合而成，首个字符可以包含星号（*），但不能以连字符（-）开头，长度为[1,256]个字符 **默认取值** 不涉及
	Dim2 *string `json:"dim.2,omitempty"`

	// **参数解释** 指标的第四层维度，目前最多支持4个维度，维度格式为dim.{i}=key,value。样例：instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d 各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/zh-cn/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档 **约束限制** 不涉及 **取值范围** dim.3=key,value，由key、value以逗号拼接组成。 key必须以字母（大写或小写）开头，后面可以跟零个或多个字母（大写或小写）、数字、下划线（_）、连字符（-），长度为[1,32]个字符 value由多个字母（大写或小写）、数字、下划线（_）、连字符（-）、点（.）、斜杠（/）、井号（#）、英文左括号（(）、英文右括号（)）组合而成，首个字符可以包含星号（*），但不能以连字符（-）开头，长度为[1,256]个字符 **默认取值** 不涉及
	Dim3 *string `json:"dim.3,omitempty"`

	// **参数解释** 聚合方式 **约束限制** 不涉及 **取值范围** 枚举值： - average 平均值 - variance 方差 - min 最小值 - max 最大值 - sum 求和值 **默认取值** 不涉及
	Filter ShowMetricDataRequestFilter `json:"filter"`

	// **参数解释** 监控数据的聚合粒度，聚合解释可查看：“[[聚合含义](https://support.huaweicloud.com/ces_faq/ces_faq_0009.html)](tag:hc)[[聚合含义](https://support.huaweicloud.com/intl/zh-cn/ces_faq/ces_faq_0009.html)](tag:hk)[[聚合含义](https://support.huaweicloud.com/eu/ces_faq/ces_faq_0009.html)](tag:hws_eu)[[聚合含义](https://docs.otc.t-systems.com/usermanual/ces/ces_faq_0009.html)](tag:dt,dt_test)[《云监控服务用户指南》中“什么是聚合？”章节](tag:ax,cmcc,ctc,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)”。 **约束限制** 不涉及 **取值范围** 枚举值： - 1 监控资源的实时数据 - 60 聚合1分钟粒度数据，表示1分钟一个数据点 - 300 聚合5分钟粒度数据，表示5分钟一个数据点 - 1200 聚合20分钟粒度数据，表示20分钟一个数据点 - 3600 聚合1小时粒度数据，表示1小时一个数据点 - 14400 聚合4小时粒度数据，表示4小时一个数据点 - 86400 聚合1天粒度数据，表示1天一个数据点 **默认取值** 不涉及
	Period ShowMetricDataRequestPeriod `json:"period"`

	// **参数解释** 查询数据起始时间，UNIX时间戳，单位毫秒。建议from的值相对于当前时间向前偏移至少1个周期。由于聚合运算的过程是将一个聚合周期范围内的数据点聚合到周期起始边界上，如果将from和to的范围设置在聚合周期内，会因为聚合未完成而造成查询数据为空，所以建议from参数相对于当前时间向前偏移至少1个周期。以5分钟聚合周期为例：假设当前时间点为10:35，10:30~10:35之间的原始数据会被聚合到10:30这个点上，所以查询5分钟数据点时from参数应为10:30或之前。云监控会根据所选择的聚合粒度向前取整from参数；如：1607146998177。 **约束限制** 不涉及 **取值范围** 毫秒级时间戳，范围为[1111111111111,9999999999999] **默认取值** 不涉及
	From int64 `json:"from"`

	// **参数解释** 查询数据截止时间UNIX时间戳，单位毫秒 **约束限制** from必须小于to **取值范围** 毫秒级时间戳，范围为[1111111111111,9999999999999] **默认取值** 不涉及
	To int64 `json:"to"`
}

func (o ShowMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricDataRequest", string(data)}, " ")
}

type ShowMetricDataRequestFilter struct {
	value string
}

type ShowMetricDataRequestFilterEnum struct {
	AVERAGE  ShowMetricDataRequestFilter
	VARIANCE ShowMetricDataRequestFilter
	MIN      ShowMetricDataRequestFilter
	MAX      ShowMetricDataRequestFilter
	SUM      ShowMetricDataRequestFilter
}

func GetShowMetricDataRequestFilterEnum() ShowMetricDataRequestFilterEnum {
	return ShowMetricDataRequestFilterEnum{
		AVERAGE: ShowMetricDataRequestFilter{
			value: "average",
		},
		VARIANCE: ShowMetricDataRequestFilter{
			value: "variance",
		},
		MIN: ShowMetricDataRequestFilter{
			value: "min",
		},
		MAX: ShowMetricDataRequestFilter{
			value: "max",
		},
		SUM: ShowMetricDataRequestFilter{
			value: "sum",
		},
	}
}

func (c ShowMetricDataRequestFilter) Value() string {
	return c.value
}

func (c ShowMetricDataRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricDataRequestFilter) UnmarshalJSON(b []byte) error {
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

type ShowMetricDataRequestPeriod struct {
	value int32
}

type ShowMetricDataRequestPeriodEnum struct {
	E_1     ShowMetricDataRequestPeriod
	E_60    ShowMetricDataRequestPeriod
	E_300   ShowMetricDataRequestPeriod
	E_1200  ShowMetricDataRequestPeriod
	E_3600  ShowMetricDataRequestPeriod
	E_14400 ShowMetricDataRequestPeriod
	E_86400 ShowMetricDataRequestPeriod
}

func GetShowMetricDataRequestPeriodEnum() ShowMetricDataRequestPeriodEnum {
	return ShowMetricDataRequestPeriodEnum{
		E_1: ShowMetricDataRequestPeriod{
			value: 1,
		}, E_60: ShowMetricDataRequestPeriod{
			value: 60,
		}, E_300: ShowMetricDataRequestPeriod{
			value: 300,
		}, E_1200: ShowMetricDataRequestPeriod{
			value: 1200,
		}, E_3600: ShowMetricDataRequestPeriod{
			value: 3600,
		}, E_14400: ShowMetricDataRequestPeriod{
			value: 14400,
		}, E_86400: ShowMetricDataRequestPeriod{
			value: 86400,
		},
	}
}

func (c ShowMetricDataRequestPeriod) Value() int32 {
	return c.value
}

func (c ShowMetricDataRequestPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricDataRequestPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
