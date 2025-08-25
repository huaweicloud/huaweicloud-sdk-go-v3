package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowVmMonitorRequest Request Object
type ShowVmMonitorRequest struct {

	// 所属空间，分别从不同的来源获取监控数据，有：ECS，DHSM
	Namespace ShowVmMonitorRequestNamespace `json:"namespace"`

	// 指标名称，有：mem_util，cpu_uttl，network_outgoing_bytes_rate_inband
	MetricName ShowVmMonitorRequestMetricName `json:"metric_name"`

	// 实例id，默认空字符串，默认查询所有实例。
	InstanceId *string `json:"instance_id,omitempty"`

	// 虚拟机id
	VsmId *string `json:"vsm_id,omitempty"`

	// 查询时间范围的起始时间，毫秒时间戳，默认0，默认查询从三天前。
	From *int64 `json:"from,omitempty"`

	// 查询时间范围的终止时间，毫秒时间戳，默认0，默认查询到当前时间。
	To *int64 `json:"to,omitempty"`

	// 统计数据周期，默认0，默认周期为5分钟
	Period *int32 `json:"period,omitempty"`

	// 统计值类型，默认min，默认查询最小值
	Filter *string `json:"filter,omitempty"`
}

func (o ShowVmMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVmMonitorRequest struct{}"
	}

	return strings.Join([]string{"ShowVmMonitorRequest", string(data)}, " ")
}

type ShowVmMonitorRequestNamespace struct {
	value string
}

type ShowVmMonitorRequestNamespaceEnum struct {
	ECS  ShowVmMonitorRequestNamespace
	DHSM ShowVmMonitorRequestNamespace
}

func GetShowVmMonitorRequestNamespaceEnum() ShowVmMonitorRequestNamespaceEnum {
	return ShowVmMonitorRequestNamespaceEnum{
		ECS: ShowVmMonitorRequestNamespace{
			value: "ECS",
		},
		DHSM: ShowVmMonitorRequestNamespace{
			value: "DHSM",
		},
	}
}

func (c ShowVmMonitorRequestNamespace) Value() string {
	return c.value
}

func (c ShowVmMonitorRequestNamespace) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVmMonitorRequestNamespace) UnmarshalJSON(b []byte) error {
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

type ShowVmMonitorRequestMetricName struct {
	value string
}

type ShowVmMonitorRequestMetricNameEnum struct {
	MEM_UTIL                           ShowVmMonitorRequestMetricName
	CPU_UTTL                           ShowVmMonitorRequestMetricName
	NETWORK_OUTGOING_BYTES_RATE_INBAND ShowVmMonitorRequestMetricName
}

func GetShowVmMonitorRequestMetricNameEnum() ShowVmMonitorRequestMetricNameEnum {
	return ShowVmMonitorRequestMetricNameEnum{
		MEM_UTIL: ShowVmMonitorRequestMetricName{
			value: "mem_util",
		},
		CPU_UTTL: ShowVmMonitorRequestMetricName{
			value: "cpu_uttl",
		},
		NETWORK_OUTGOING_BYTES_RATE_INBAND: ShowVmMonitorRequestMetricName{
			value: "network_outgoing_bytes_rate_inband",
		},
	}
}

func (c ShowVmMonitorRequestMetricName) Value() string {
	return c.value
}

func (c ShowVmMonitorRequestMetricName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVmMonitorRequestMetricName) UnmarshalJSON(b []byte) error {
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
