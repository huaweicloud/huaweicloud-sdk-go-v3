package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListFunctionStatisticsRequest struct {

	// 函数的URN（Uniform Resource Name），唯一标识函数。
	FuncUrn string `json:"func_urn"`

	// 获取最近多少分钟内函数执行的指标。
	Period ListFunctionStatisticsRequestPeriod `json:"period"`
}

func (o ListFunctionStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionStatisticsRequest", string(data)}, " ")
}

type ListFunctionStatisticsRequestPeriod struct {
	value string
}

type ListFunctionStatisticsRequestPeriodEnum struct {
	E_5  ListFunctionStatisticsRequestPeriod
	E_15 ListFunctionStatisticsRequestPeriod
	E_60 ListFunctionStatisticsRequestPeriod
}

func GetListFunctionStatisticsRequestPeriodEnum() ListFunctionStatisticsRequestPeriodEnum {
	return ListFunctionStatisticsRequestPeriodEnum{
		E_5: ListFunctionStatisticsRequestPeriod{
			value: "5",
		},
		E_15: ListFunctionStatisticsRequestPeriod{
			value: "15",
		},
		E_60: ListFunctionStatisticsRequestPeriod{
			value: "60",
		},
	}
}

func (c ListFunctionStatisticsRequestPeriod) Value() string {
	return c.value
}

func (c ListFunctionStatisticsRequestPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFunctionStatisticsRequestPeriod) UnmarshalJSON(b []byte) error {
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
