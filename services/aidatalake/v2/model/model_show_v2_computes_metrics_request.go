package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowV2ComputesMetricsRequest Request Object
type ShowV2ComputesMetricsRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：查询时间步长。 **约束限制**：不涉及。 **取值范围**：单位秒。   - 60：每分钟一个点。   - 3600：每1小时一个点。   - 86400：每天一个点。 **默认取值**：60。
	Step ShowV2ComputesMetricsRequestStep `json:"step"`

	// **参数解释**：指标名称。 **约束限制**：不涉及。 **取值范围**：   - CPU：资源池CPU使用率。   - MEM：资源池内存使用率。   - NPU：资源池NPU使用率。   - GPU：资源池GPU使用率。 **默认取值**：不涉及。
	MetricName ShowV2ComputesMetricsRequestMetricName `json:"metric_name"`

	// **参数解释**：开始时间。 **约束限制**：不涉及。 **取值范围**：时间戳，单位秒。 **默认取值**：不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**：结束时间。 **约束限制**：不涉及。 **取值范围**：时间戳，单位秒。 **默认取值**：不涉及。
	EndTime int64 `json:"end_time"`

	// **参数解释**：资源池ids，多个用英文逗号分割。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：1~1110个英文字符。
	ComputeIds string `json:"compute_ids"`
}

func (o ShowV2ComputesMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2ComputesMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowV2ComputesMetricsRequest", string(data)}, " ")
}

type ShowV2ComputesMetricsRequestStep struct {
	value int32
}

type ShowV2ComputesMetricsRequestStepEnum struct {
	E_60    ShowV2ComputesMetricsRequestStep
	E_3600  ShowV2ComputesMetricsRequestStep
	E_86400 ShowV2ComputesMetricsRequestStep
}

func GetShowV2ComputesMetricsRequestStepEnum() ShowV2ComputesMetricsRequestStepEnum {
	return ShowV2ComputesMetricsRequestStepEnum{
		E_60: ShowV2ComputesMetricsRequestStep{
			value: 60,
		}, E_3600: ShowV2ComputesMetricsRequestStep{
			value: 3600,
		}, E_86400: ShowV2ComputesMetricsRequestStep{
			value: 86400,
		},
	}
}

func (c ShowV2ComputesMetricsRequestStep) Value() int32 {
	return c.value
}

func (c ShowV2ComputesMetricsRequestStep) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowV2ComputesMetricsRequestStep) UnmarshalJSON(b []byte) error {
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

type ShowV2ComputesMetricsRequestMetricName struct {
	value string
}

type ShowV2ComputesMetricsRequestMetricNameEnum struct {
	CPU ShowV2ComputesMetricsRequestMetricName
	NPU ShowV2ComputesMetricsRequestMetricName
	GPU ShowV2ComputesMetricsRequestMetricName
	MEM ShowV2ComputesMetricsRequestMetricName
}

func GetShowV2ComputesMetricsRequestMetricNameEnum() ShowV2ComputesMetricsRequestMetricNameEnum {
	return ShowV2ComputesMetricsRequestMetricNameEnum{
		CPU: ShowV2ComputesMetricsRequestMetricName{
			value: "CPU",
		},
		NPU: ShowV2ComputesMetricsRequestMetricName{
			value: "NPU",
		},
		GPU: ShowV2ComputesMetricsRequestMetricName{
			value: "GPU",
		},
		MEM: ShowV2ComputesMetricsRequestMetricName{
			value: "MEM",
		},
	}
}

func (c ShowV2ComputesMetricsRequestMetricName) Value() string {
	return c.value
}

func (c ShowV2ComputesMetricsRequestMetricName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowV2ComputesMetricsRequestMetricName) UnmarshalJSON(b []byte) error {
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
