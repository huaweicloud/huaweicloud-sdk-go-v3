package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowMetricResultResponseBody struct {

	// 指标ID
	MetricId string `json:"metric_id"`

	Result *ShowMetricResultResponseBodyResult `json:"result"`

	// 指标显示格式，根据不同指标固定返回。
	MetricFormat *[]MetricFormat `json:"metric_format,omitempty"`

	// 结果日志信息
	LogMsg *string `json:"log_msg,omitempty"`

	// 查询结果状态，SUCCESS：查询成功，FAILED：查询失败，FALLBACK：使用默认值
	Status *ShowMetricResultResponseBodyStatus `json:"status,omitempty"`
}

func (o ShowMetricResultResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricResultResponseBody struct{}"
	}

	return strings.Join([]string{"ShowMetricResultResponseBody", string(data)}, " ")
}

type ShowMetricResultResponseBodyStatus struct {
	value string
}

type ShowMetricResultResponseBodyStatusEnum struct {
	SUCCESS  ShowMetricResultResponseBodyStatus
	FAILED   ShowMetricResultResponseBodyStatus
	FALLBACK ShowMetricResultResponseBodyStatus
}

func GetShowMetricResultResponseBodyStatusEnum() ShowMetricResultResponseBodyStatusEnum {
	return ShowMetricResultResponseBodyStatusEnum{
		SUCCESS: ShowMetricResultResponseBodyStatus{
			value: "SUCCESS",
		},
		FAILED: ShowMetricResultResponseBodyStatus{
			value: "FAILED",
		},
		FALLBACK: ShowMetricResultResponseBodyStatus{
			value: "FALLBACK",
		},
	}
}

func (c ShowMetricResultResponseBodyStatus) Value() string {
	return c.value
}

func (c ShowMetricResultResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricResultResponseBodyStatus) UnmarshalJSON(b []byte) error {
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
