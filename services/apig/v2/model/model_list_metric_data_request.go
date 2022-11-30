package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListMetricDataRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 指标维度 - inbound_eip：入口公网带宽，仅ELB类型实例支持 - outbound_eip：出口公网带宽
	Dim ListMetricDataRequestDim `json:"dim"`

	// 指标名称 - upstream_bandwidth：出网带宽 - downstream_bandwidth：入网带宽 - upstream_bandwidth_usage：出网带宽使用率 - downstream_bandwidth_usage：入网带宽使用率 - up_stream：出网流量 - down_stream：入网流量
	MetricName ListMetricDataRequestMetricName `json:"metric_name"`

	// 查询数据起始时间，UNIX时间戳，单位毫秒。
	From string `json:"from"`

	// 查询数据截止时间UNIX时间戳，单位毫秒。from必须小于to。
	To string `json:"to"`

	// 监控数据粒度。 - 1：实时数据 - 300：5分钟粒度 - 1200：20分钟粒度 - 3600：1小时粒度 - 14400：4小时粒度 - 86400：1天粒度
	Period ListMetricDataRequestPeriod `json:"period"`

	// 数据聚合方式。 - average：聚合周期内指标数据的平均值。 - max：聚合周期内指标数据的最大值。 - min：聚合周期内指标数据的最小值。 - sum：聚合周期内指标数据的求和值。 - variance：聚合周期内指标数据的方差。
	Filter ListMetricDataRequestFilter `json:"filter"`
}

func (o ListMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ListMetricDataRequest", string(data)}, " ")
}

type ListMetricDataRequestDim struct {
	value string
}

type ListMetricDataRequestDimEnum struct {
	INBOUND  ListMetricDataRequestDim
	OUTBOUND ListMetricDataRequestDim
}

func GetListMetricDataRequestDimEnum() ListMetricDataRequestDimEnum {
	return ListMetricDataRequestDimEnum{
		INBOUND: ListMetricDataRequestDim{
			value: "inbound",
		},
		OUTBOUND: ListMetricDataRequestDim{
			value: "outbound",
		},
	}
}

func (c ListMetricDataRequestDim) Value() string {
	return c.value
}

func (c ListMetricDataRequestDim) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricDataRequestDim) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListMetricDataRequestMetricName struct {
	value string
}

type ListMetricDataRequestMetricNameEnum struct {
	UPSTREAM_BANDWIDTH         ListMetricDataRequestMetricName
	DOWNSTREAM_BANDWIDTH       ListMetricDataRequestMetricName
	UPSTREAM_BANDWIDTH_USAGE   ListMetricDataRequestMetricName
	DOWNSTREAM_BANDWIDTH_USAGE ListMetricDataRequestMetricName
	UP_STREAM                  ListMetricDataRequestMetricName
	DOWN_STREAM                ListMetricDataRequestMetricName
}

func GetListMetricDataRequestMetricNameEnum() ListMetricDataRequestMetricNameEnum {
	return ListMetricDataRequestMetricNameEnum{
		UPSTREAM_BANDWIDTH: ListMetricDataRequestMetricName{
			value: "upstream_bandwidth",
		},
		DOWNSTREAM_BANDWIDTH: ListMetricDataRequestMetricName{
			value: "downstream_bandwidth",
		},
		UPSTREAM_BANDWIDTH_USAGE: ListMetricDataRequestMetricName{
			value: "upstream_bandwidth_usage",
		},
		DOWNSTREAM_BANDWIDTH_USAGE: ListMetricDataRequestMetricName{
			value: "downstream_bandwidth_usage",
		},
		UP_STREAM: ListMetricDataRequestMetricName{
			value: "up_stream",
		},
		DOWN_STREAM: ListMetricDataRequestMetricName{
			value: "down_stream",
		},
	}
}

func (c ListMetricDataRequestMetricName) Value() string {
	return c.value
}

func (c ListMetricDataRequestMetricName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricDataRequestMetricName) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListMetricDataRequestPeriod struct {
	value int32
}

type ListMetricDataRequestPeriodEnum struct {
	E_1     ListMetricDataRequestPeriod
	E_300   ListMetricDataRequestPeriod
	E_1200  ListMetricDataRequestPeriod
	E_3600  ListMetricDataRequestPeriod
	E_14400 ListMetricDataRequestPeriod
	E_86400 ListMetricDataRequestPeriod
}

func GetListMetricDataRequestPeriodEnum() ListMetricDataRequestPeriodEnum {
	return ListMetricDataRequestPeriodEnum{
		E_1: ListMetricDataRequestPeriod{
			value: 1,
		}, E_300: ListMetricDataRequestPeriod{
			value: 300,
		}, E_1200: ListMetricDataRequestPeriod{
			value: 1200,
		}, E_3600: ListMetricDataRequestPeriod{
			value: 3600,
		}, E_14400: ListMetricDataRequestPeriod{
			value: 14400,
		}, E_86400: ListMetricDataRequestPeriod{
			value: 86400,
		},
	}
}

func (c ListMetricDataRequestPeriod) Value() int32 {
	return c.value
}

func (c ListMetricDataRequestPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricDataRequestPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListMetricDataRequestFilter struct {
	value string
}

type ListMetricDataRequestFilterEnum struct {
	AVERAGE  ListMetricDataRequestFilter
	MAX      ListMetricDataRequestFilter
	MIN      ListMetricDataRequestFilter
	SUM      ListMetricDataRequestFilter
	VARIANCE ListMetricDataRequestFilter
}

func GetListMetricDataRequestFilterEnum() ListMetricDataRequestFilterEnum {
	return ListMetricDataRequestFilterEnum{
		AVERAGE: ListMetricDataRequestFilter{
			value: "average",
		},
		MAX: ListMetricDataRequestFilter{
			value: "max",
		},
		MIN: ListMetricDataRequestFilter{
			value: "min",
		},
		SUM: ListMetricDataRequestFilter{
			value: "sum",
		},
		VARIANCE: ListMetricDataRequestFilter{
			value: "variance",
		},
	}
}

func (c ListMetricDataRequestFilter) Value() string {
	return c.value
}

func (c ListMetricDataRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricDataRequestFilter) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
