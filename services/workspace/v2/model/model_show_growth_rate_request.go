package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowGrowthRateRequest Request Object
type ShowGrowthRateRequest struct {

	// 环比周期 | DAY - 天 MONTH - 月
	GrowPeriod *ShowGrowthRateRequestGrowPeriod `json:"grow_period,omitempty"`

	// 指标名称
	MetricName string `json:"metric_name"`

	// 指标维度 | 目前最大支持3个维度，必须从0开始；维度格式为dim.{i}=key,value，key的最大长度32，value的最大长度为256。 单维度：dim.0=instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d 多维度：dim.0=key,value&dim.1=key,value
	Dim *string `json:"dim,omitempty"`
}

func (o ShowGrowthRateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGrowthRateRequest struct{}"
	}

	return strings.Join([]string{"ShowGrowthRateRequest", string(data)}, " ")
}

type ShowGrowthRateRequestGrowPeriod struct {
	value string
}

type ShowGrowthRateRequestGrowPeriodEnum struct {
	DAY   ShowGrowthRateRequestGrowPeriod
	MONTH ShowGrowthRateRequestGrowPeriod
}

func GetShowGrowthRateRequestGrowPeriodEnum() ShowGrowthRateRequestGrowPeriodEnum {
	return ShowGrowthRateRequestGrowPeriodEnum{
		DAY: ShowGrowthRateRequestGrowPeriod{
			value: "DAY",
		},
		MONTH: ShowGrowthRateRequestGrowPeriod{
			value: "MONTH",
		},
	}
}

func (c ShowGrowthRateRequestGrowPeriod) Value() string {
	return c.value
}

func (c ShowGrowthRateRequestGrowPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGrowthRateRequestGrowPeriod) UnmarshalJSON(b []byte) error {
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
