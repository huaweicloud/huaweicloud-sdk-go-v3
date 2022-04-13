package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type CollectSessionStatsRequest struct {
	// qabot编号，UUID格式。

	QabotId string `json:"qabot_id"`
	// 查询的起始时间，long，UTC时间，默认值为0。

	StartTime *string `json:"start_time,omitempty"`
	// 查询的结束时间，long，UTC时间，默认值为当前时间的毫秒数。

	EndTime *string `json:"end_time,omitempty"`
	// 统计周期目前支持month,week,day。

	Interval *CollectSessionStatsRequestInterval `json:"interval,omitempty"`
	// 请求所在时区，例如：中国东八区为\"+08:00\"；美国西五区为\"-05:00\"；默认为\"UTC\"。

	TimeZone *string `json:"time_zone,omitempty"`
}

func (o CollectSessionStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectSessionStatsRequest struct{}"
	}

	return strings.Join([]string{"CollectSessionStatsRequest", string(data)}, " ")
}

type CollectSessionStatsRequestInterval struct {
	value string
}

type CollectSessionStatsRequestIntervalEnum struct {
	MONTH CollectSessionStatsRequestInterval
	WEEK  CollectSessionStatsRequestInterval
	DAY   CollectSessionStatsRequestInterval
}

func GetCollectSessionStatsRequestIntervalEnum() CollectSessionStatsRequestIntervalEnum {
	return CollectSessionStatsRequestIntervalEnum{
		MONTH: CollectSessionStatsRequestInterval{
			value: "month",
		},
		WEEK: CollectSessionStatsRequestInterval{
			value: "week",
		},
		DAY: CollectSessionStatsRequestInterval{
			value: "day",
		},
	}
}

func (c CollectSessionStatsRequestInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectSessionStatsRequestInterval) UnmarshalJSON(b []byte) error {
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
