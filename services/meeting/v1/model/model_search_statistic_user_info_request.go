package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SearchStatisticUserInfoRequest struct {
	// 查询偏移量。 * 取值：大于等于0，默认值为0。 * 小于最小值0时，系统设置为0。 * 大于等于最大条目数量，则返回最后一页数据，页数根据总条目数和limit计算得出。

	Offset *int32 `json:"offset,omitempty"`
	// 查询的条目数量。 * 取值：1-500，默认值为20。 * 小于最小值1时，系统设置为1。 * 大于最大值500时，系统设置为500。

	Limit *int32 `json:"limit,omitempty"`
	// 查询时间维度，取值： * D: 按日查询 * M: 按月查询。

	TimeUnit SearchStatisticUserInfoRequestTimeUnit `json:"timeUnit"`
	// 查询时间范围的开始时间，格式根据timeUnit的取值而定。 * timeUnit = D，格式：yyyy-MM-dd，此情况下startTime与endTime间隔最多31日。 * timeUnit = M，格式：yyyy-MM，此情况下startTime与endTime间隔最多12个月。

	StartTime string `json:"startTime"`
	// 查询时间范围的结束时间，格式根据timeUnit的取值而定。 * timeUnit = D，格式：yyyy-MM-dd，此情况下startTime与endTime间隔最多31日。 * timeUnit = M，格式：yyyy-MM，此情况下startTime与endTime间隔最多12个月。

	EndTime string `json:"endTime"`
	// 查询分类，取值： * user_login_info: 用户登录数据 * user_activate_info: 用户激活数据 * user_login_device_info: 用户登录设备数据

	Category SearchStatisticUserInfoRequestCategory `json:"category"`
}

func (o SearchStatisticUserInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchStatisticUserInfoRequest struct{}"
	}

	return strings.Join([]string{"SearchStatisticUserInfoRequest", string(data)}, " ")
}

type SearchStatisticUserInfoRequestTimeUnit struct {
	value string
}

type SearchStatisticUserInfoRequestTimeUnitEnum struct {
	D SearchStatisticUserInfoRequestTimeUnit
	M SearchStatisticUserInfoRequestTimeUnit
}

func GetSearchStatisticUserInfoRequestTimeUnitEnum() SearchStatisticUserInfoRequestTimeUnitEnum {
	return SearchStatisticUserInfoRequestTimeUnitEnum{
		D: SearchStatisticUserInfoRequestTimeUnit{
			value: "D",
		},
		M: SearchStatisticUserInfoRequestTimeUnit{
			value: "M",
		},
	}
}

func (c SearchStatisticUserInfoRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchStatisticUserInfoRequestTimeUnit) UnmarshalJSON(b []byte) error {
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

type SearchStatisticUserInfoRequestCategory struct {
	value string
}

type SearchStatisticUserInfoRequestCategoryEnum struct {
	USER_LOGIN_INFO        SearchStatisticUserInfoRequestCategory
	USER_ACTIVATE_INFO     SearchStatisticUserInfoRequestCategory
	USER_LOGIN_DEVICE_INFO SearchStatisticUserInfoRequestCategory
}

func GetSearchStatisticUserInfoRequestCategoryEnum() SearchStatisticUserInfoRequestCategoryEnum {
	return SearchStatisticUserInfoRequestCategoryEnum{
		USER_LOGIN_INFO: SearchStatisticUserInfoRequestCategory{
			value: "user_login_info",
		},
		USER_ACTIVATE_INFO: SearchStatisticUserInfoRequestCategory{
			value: "user_activate_info",
		},
		USER_LOGIN_DEVICE_INFO: SearchStatisticUserInfoRequestCategory{
			value: "user_login_device_info",
		},
	}
}

func (c SearchStatisticUserInfoRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchStatisticUserInfoRequestCategory) UnmarshalJSON(b []byte) error {
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
