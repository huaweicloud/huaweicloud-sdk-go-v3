package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SearchStatisticConferenceInfoRequest struct {

	// 查询偏移量。 * 取值：大于等于0，默认值为0。 * 小于最小值0时，系统设置为0。 * 大于等于最大条目数量，则返回最后一页数据，页数根据总条目数和limit计算得出。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询的条目数量。 * 取值：1-500，默认值为20。 * 小于最小值1时，系统设置为1。 * 大于最大值500时，系统设置为500。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询时间维度，取值： * D: 按日查询 * M: 按月查询。
	TimeUnit SearchStatisticConferenceInfoRequestTimeUnit `json:"timeUnit" xml:"timeUnit"`

	// 查询时间范围的开始时间，格式根据timeUnit的取值而定。 * timeUnit = D，格式：yyyy-MM-dd，此情况下startTime与endTime间隔最多31日。 * timeUnit = M，格式：yyyy-MM，此情况下startTime与endTime间隔最多12个月。
	StartTime string `json:"startTime" xml:"startTime"`

	// 查询时间范围的结束时间，格式根据timeUnit的取值而定。 * timeUnit = D，格式：yyyy-MM-dd，此情况下startTime与endTime间隔最多31日。 * timeUnit = M，格式：yyyy-MM，此情况下startTime与endTime间隔最多12个月。
	EndTime string `json:"endTime" xml:"endTime"`

	// 查询分类，取值： * conference_info: 会议总体数据 * conference_hourly_info: 单日内会议总体数据 * category = conference_hourly_info的情况，timeUnit只能取值'D'，且startTime与endTime必须为同一天
	Category SearchStatisticConferenceInfoRequestCategory `json:"category" xml:"category"`
}

func (o SearchStatisticConferenceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchStatisticConferenceInfoRequest struct{}"
	}

	return strings.Join([]string{"SearchStatisticConferenceInfoRequest", string(data)}, " ")
}

type SearchStatisticConferenceInfoRequestTimeUnit struct {
	value string
}

type SearchStatisticConferenceInfoRequestTimeUnitEnum struct {
	D SearchStatisticConferenceInfoRequestTimeUnit
	M SearchStatisticConferenceInfoRequestTimeUnit
}

func GetSearchStatisticConferenceInfoRequestTimeUnitEnum() SearchStatisticConferenceInfoRequestTimeUnitEnum {
	return SearchStatisticConferenceInfoRequestTimeUnitEnum{
		D: SearchStatisticConferenceInfoRequestTimeUnit{
			value: "D",
		},
		M: SearchStatisticConferenceInfoRequestTimeUnit{
			value: "M",
		},
	}
}

func (c SearchStatisticConferenceInfoRequestTimeUnit) Value() string {
	return c.value
}

func (c SearchStatisticConferenceInfoRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchStatisticConferenceInfoRequestTimeUnit) UnmarshalJSON(b []byte) error {
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

type SearchStatisticConferenceInfoRequestCategory struct {
	value string
}

type SearchStatisticConferenceInfoRequestCategoryEnum struct {
	CONFERENCE_INFO        SearchStatisticConferenceInfoRequestCategory
	CONFERENCE_HOURLY_INFO SearchStatisticConferenceInfoRequestCategory
}

func GetSearchStatisticConferenceInfoRequestCategoryEnum() SearchStatisticConferenceInfoRequestCategoryEnum {
	return SearchStatisticConferenceInfoRequestCategoryEnum{
		CONFERENCE_INFO: SearchStatisticConferenceInfoRequestCategory{
			value: "conference_info",
		},
		CONFERENCE_HOURLY_INFO: SearchStatisticConferenceInfoRequestCategory{
			value: "conference_hourly_info",
		},
	}
}

func (c SearchStatisticConferenceInfoRequestCategory) Value() string {
	return c.value
}

func (c SearchStatisticConferenceInfoRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchStatisticConferenceInfoRequestCategory) UnmarshalJSON(b []byte) error {
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
