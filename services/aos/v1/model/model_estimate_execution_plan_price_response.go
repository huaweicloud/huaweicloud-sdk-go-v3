package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type EstimateExecutionPlanPriceResponse struct {

	// 币种，枚举值   * `CNY` - 元，中国站返回的币种   * `USD` - 美元，国际站返回的币种
	Currency *EstimateExecutionPlanPriceResponseCurrency `json:"currency,omitempty"`

	// 执行计划中所有资源的询价结果
	Items          *[]ItemsResponse `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o EstimateExecutionPlanPriceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateExecutionPlanPriceResponse struct{}"
	}

	return strings.Join([]string{"EstimateExecutionPlanPriceResponse", string(data)}, " ")
}

type EstimateExecutionPlanPriceResponseCurrency struct {
	value string
}

type EstimateExecutionPlanPriceResponseCurrencyEnum struct {
	CNY EstimateExecutionPlanPriceResponseCurrency
	USD EstimateExecutionPlanPriceResponseCurrency
}

func GetEstimateExecutionPlanPriceResponseCurrencyEnum() EstimateExecutionPlanPriceResponseCurrencyEnum {
	return EstimateExecutionPlanPriceResponseCurrencyEnum{
		CNY: EstimateExecutionPlanPriceResponseCurrency{
			value: "CNY",
		},
		USD: EstimateExecutionPlanPriceResponseCurrency{
			value: "USD",
		},
	}
}

func (c EstimateExecutionPlanPriceResponseCurrency) Value() string {
	return c.value
}

func (c EstimateExecutionPlanPriceResponseCurrency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EstimateExecutionPlanPriceResponseCurrency) UnmarshalJSON(b []byte) error {
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
