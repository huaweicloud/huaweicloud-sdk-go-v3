package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SearchStatisticResourceInfoRequest struct {

	// 查询偏移量。 * 取值：大于等于0，默认值为0。 * 小于最小值0时，系统设置为0。 * 大于等于最大条目数量，则返回最后一页数据，页数根据总条目数和limit计算得出。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的条目数量。 * 取值：1-500，默认值为20。 * 小于最小值1时，系统设置为1。 * 大于最大值500时，系统设置为500。
	Limit *int32 `json:"limit,omitempty"`

	// 查询时间维度，取值： * D: 按日查询 * M: 按月查询。
	TimeUnit SearchStatisticResourceInfoRequestTimeUnit `json:"timeUnit"`

	// 查询时间范围的开始时间，格式根据timeUnit的取值而定。 * timeUnit = D，格式：yyyy-MM-dd，此情况下startTime与endTime间隔最多31日。 * timeUnit = M，格式：yyyy-MM，此情况下startTime与endTime间隔最多12个月。
	StartTime string `json:"startTime"`

	// 查询时间范围的结束时间，格式根据timeUnit的取值而定。 * timeUnit = D，格式：yyyy-MM-dd，此情况下startTime与endTime间隔最多31日。 * timeUnit = M，格式：yyyy-MM，此情况下startTime与endTime间隔最多12个月。
	EndTime string `json:"endTime"`

	// 查询分类，取值： * used_vmr_info: 已购VMR资源使用统计数据 * used_live_info: 已购直播端口资源使用统计数据 * used_record_info: 已购录播资源使用统计数据 * used_pstn_info: 已购电话外呼资源使用统计数据
	Category SearchStatisticResourceInfoRequestCategory `json:"category"`
}

func (o SearchStatisticResourceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchStatisticResourceInfoRequest struct{}"
	}

	return strings.Join([]string{"SearchStatisticResourceInfoRequest", string(data)}, " ")
}

type SearchStatisticResourceInfoRequestTimeUnit struct {
	value string
}

type SearchStatisticResourceInfoRequestTimeUnitEnum struct {
	D SearchStatisticResourceInfoRequestTimeUnit
	M SearchStatisticResourceInfoRequestTimeUnit
}

func GetSearchStatisticResourceInfoRequestTimeUnitEnum() SearchStatisticResourceInfoRequestTimeUnitEnum {
	return SearchStatisticResourceInfoRequestTimeUnitEnum{
		D: SearchStatisticResourceInfoRequestTimeUnit{
			value: "D",
		},
		M: SearchStatisticResourceInfoRequestTimeUnit{
			value: "M",
		},
	}
}

func (c SearchStatisticResourceInfoRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchStatisticResourceInfoRequestTimeUnit) UnmarshalJSON(b []byte) error {
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

type SearchStatisticResourceInfoRequestCategory struct {
	value string
}

type SearchStatisticResourceInfoRequestCategoryEnum struct {
	USED_VMR_INFO    SearchStatisticResourceInfoRequestCategory
	USED_LIVE_INFO   SearchStatisticResourceInfoRequestCategory
	USED_RECORD_INFO SearchStatisticResourceInfoRequestCategory
	USED_PSTN_INFO   SearchStatisticResourceInfoRequestCategory
}

func GetSearchStatisticResourceInfoRequestCategoryEnum() SearchStatisticResourceInfoRequestCategoryEnum {
	return SearchStatisticResourceInfoRequestCategoryEnum{
		USED_VMR_INFO: SearchStatisticResourceInfoRequestCategory{
			value: "used_vmr_info",
		},
		USED_LIVE_INFO: SearchStatisticResourceInfoRequestCategory{
			value: "used_live_info",
		},
		USED_RECORD_INFO: SearchStatisticResourceInfoRequestCategory{
			value: "used_record_info",
		},
		USED_PSTN_INFO: SearchStatisticResourceInfoRequestCategory{
			value: "used_pstn_info",
		},
	}
}

func (c SearchStatisticResourceInfoRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchStatisticResourceInfoRequestCategory) UnmarshalJSON(b []byte) error {
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
