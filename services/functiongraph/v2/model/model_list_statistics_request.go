package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListStatisticsRequest struct {
	// 参数过滤器。

	Filter ListStatisticsRequestFilter `json:"filter"`
	// 时间段单位为分钟，与filter参数配合使用。

	Period *string `json:"period,omitempty"`
}

func (o ListStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListStatisticsRequest", string(data)}, " ")
}

type ListStatisticsRequestFilter struct {
	value string
}

type ListStatisticsRequestFilterEnum struct {
	MONITOR_DATA   ListStatisticsRequestFilter
	MONTHLY_REPORT ListStatisticsRequestFilter
}

func GetListStatisticsRequestFilterEnum() ListStatisticsRequestFilterEnum {
	return ListStatisticsRequestFilterEnum{
		MONITOR_DATA: ListStatisticsRequestFilter{
			value: "monitor_data",
		},
		MONTHLY_REPORT: ListStatisticsRequestFilter{
			value: "monthly_report",
		},
	}
}

func (c ListStatisticsRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStatisticsRequestFilter) UnmarshalJSON(b []byte) error {
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
