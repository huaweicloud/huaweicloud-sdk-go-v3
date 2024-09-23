package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMetricsRequest Request Object
type ListMetricsRequest struct {

	// 开始时间, UTC时间, 格式为：yyyy-MM-ddTHH:mm:ssZ
	StartTime string `json:"start_time"`

	// 结束时间 ,UTC时间, 格式为：yyyy-MM-ddTHH:mm:ssZ
	EndTime string `json:"end_time"`

	// 指标维度 | 目前最大支持3个维度，必须从0开始；维度格式为dim.{i}=key,value，key的最大长度32，value的最大长度为256。 单维度：dim.0=instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d 多维度：dim.0=key,value&dim.1=key,value
	Dim *string `json:"dim,omitempty"`

	// 指标名称，支持的指标名称参考[Workspace运维监控指标](https://support.huaweicloud.com/usermanual-workspace/workspace_06_1032.html)
	MetricNames []string `json:"metric_names"`

	// 数据周期 | DAY - 天级数据 HOUR - 小时级数据
	Period *ListMetricsRequestPeriod `json:"period,omitempty"`
}

func (o ListMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListMetricsRequest", string(data)}, " ")
}

type ListMetricsRequestPeriod struct {
	value string
}

type ListMetricsRequestPeriodEnum struct {
	DAY  ListMetricsRequestPeriod
	HOUR ListMetricsRequestPeriod
}

func GetListMetricsRequestPeriodEnum() ListMetricsRequestPeriodEnum {
	return ListMetricsRequestPeriodEnum{
		DAY: ListMetricsRequestPeriod{
			value: "DAY",
		},
		HOUR: ListMetricsRequestPeriod{
			value: "HOUR",
		},
	}
}

func (c ListMetricsRequestPeriod) Value() string {
	return c.value
}

func (c ListMetricsRequestPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricsRequestPeriod) UnmarshalJSON(b []byte) error {
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
