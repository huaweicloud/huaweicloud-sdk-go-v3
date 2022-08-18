package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListStatisticsRequest struct {

	// 参数过滤器。 monitor_data: 查询统计信息。 monthly_report：查询月度统计信息。
	Filter ListStatisticsRequestFilter `json:"filter"`

	// 时间段单位为分钟，与filter参数metric配合使用。
	Period *string `json:"period,omitempty"`

	// 月度统计的维度，filter参数取值为monthly_report时才生效。 当取值不在以上范围时，默认取\"0\"。 - \"0\": 表示统计本月。 - \"1\": 表示统计上月。 - \"2\": 表示统计最近三个月。 - \"3\": 表示统计最近六个月。
	Option *string `json:"option,omitempty"`
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

func (c ListStatisticsRequestFilter) Value() string {
	return c.value
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
