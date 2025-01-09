package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ListFutureExecutionsReq 获取未来执行的具体时间列表请求体。
type ListFutureExecutionsReq struct {

	// 执行周期类型，可选值为： - FIXED_TIME：指定时间。 - DAY：按天。 - WEEK：按周。 - MONTH：按月。
	ScheduledType *ListFutureExecutionsReqScheduledType `json:"scheduled_type,omitempty"`

	// 周期按天时：按x天间隔执行。
	DayInterval *int32 `json:"day_interval,omitempty"`

	// 周期按周时：取值1~7，英文逗号分隔，如1,2,7。
	WeekList *string `json:"week_list,omitempty"`

	// 周期按月时：取值1~12，英文逗号分隔，如1,3,12。
	MonthList *string `json:"month_list,omitempty"`

	// 周期按月时：取值1~31及L(代表当月最后一天)，英文逗号分隔，如1,2,28,L。
	DateList *string `json:"date_list,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 周期指定时间时：表示指定的日期，格式为yyyy-MM-dd。
	ScheduledDate *string `json:"scheduled_date,omitempty"`

	// 指定的执行时间点，格式为HH:mm:ss。
	ScheduledTime *string `json:"scheduled_time,omitempty"`

	// 到期时间。
	ExpireTime *sdktime.SdkTime `json:"expire_time,omitempty"`
}

func (o ListFutureExecutionsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFutureExecutionsReq struct{}"
	}

	return strings.Join([]string{"ListFutureExecutionsReq", string(data)}, " ")
}

type ListFutureExecutionsReqScheduledType struct {
	value string
}

type ListFutureExecutionsReqScheduledTypeEnum struct {
	FIXED_TIME ListFutureExecutionsReqScheduledType
	DAY        ListFutureExecutionsReqScheduledType
	WEEK       ListFutureExecutionsReqScheduledType
	MONTH      ListFutureExecutionsReqScheduledType
}

func GetListFutureExecutionsReqScheduledTypeEnum() ListFutureExecutionsReqScheduledTypeEnum {
	return ListFutureExecutionsReqScheduledTypeEnum{
		FIXED_TIME: ListFutureExecutionsReqScheduledType{
			value: "FIXED_TIME",
		},
		DAY: ListFutureExecutionsReqScheduledType{
			value: "DAY",
		},
		WEEK: ListFutureExecutionsReqScheduledType{
			value: "WEEK",
		},
		MONTH: ListFutureExecutionsReqScheduledType{
			value: "MONTH",
		},
	}
}

func (c ListFutureExecutionsReqScheduledType) Value() string {
	return c.value
}

func (c ListFutureExecutionsReqScheduledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFutureExecutionsReqScheduledType) UnmarshalJSON(b []byte) error {
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
