package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectReplyRatesRequest Request Object
type CollectReplyRatesRequest struct {

	// qabot编号，UUID格式。
	QabotId string `json:"qabot_id"`

	// 查询的起始时间，long，UTC时间，默认值为0。
	StartTime *string `json:"start_time,omitempty"`

	// 查询的结束时间，long，UTC时间，默认值为当前时间的毫秒数。
	EndTime *string `json:"end_time,omitempty"`

	// 统计周期目前支持month,week,day。
	Interval *CollectReplyRatesRequestInterval `json:"interval,omitempty"`

	// 请求所在时区，例如：中国东八区为\"+08:00\"；美国西五区为\"-05:00\"；默认为\"UTC\"。
	TimeZone *string `json:"time_zone,omitempty"`

	// 所属领域。
	Domain *string `json:"domain,omitempty"`
}

func (o CollectReplyRatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectReplyRatesRequest struct{}"
	}

	return strings.Join([]string{"CollectReplyRatesRequest", string(data)}, " ")
}

type CollectReplyRatesRequestInterval struct {
	value string
}

type CollectReplyRatesRequestIntervalEnum struct {
	MONTH CollectReplyRatesRequestInterval
	WEEK  CollectReplyRatesRequestInterval
	DAY   CollectReplyRatesRequestInterval
}

func GetCollectReplyRatesRequestIntervalEnum() CollectReplyRatesRequestIntervalEnum {
	return CollectReplyRatesRequestIntervalEnum{
		MONTH: CollectReplyRatesRequestInterval{
			value: "month",
		},
		WEEK: CollectReplyRatesRequestInterval{
			value: "week",
		},
		DAY: CollectReplyRatesRequestInterval{
			value: "day",
		},
	}
}

func (c CollectReplyRatesRequestInterval) Value() string {
	return c.value
}

func (c CollectReplyRatesRequestInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectReplyRatesRequestInterval) UnmarshalJSON(b []byte) error {
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
