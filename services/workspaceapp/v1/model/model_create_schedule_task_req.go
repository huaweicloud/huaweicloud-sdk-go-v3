package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateScheduleTaskReq 新增定时任务。
type CreateScheduleTaskReq struct {

	// 执行周期类型，可选值为： - FIXED_TIME：指定时间。 - DAY：按天。 - WEEK：按周。 - MONTH：按月。
	ScheduledType *CreateScheduleTaskReqScheduledType `json:"scheduled_type,omitempty"`

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

func (o CreateScheduleTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleTaskReq struct{}"
	}

	return strings.Join([]string{"CreateScheduleTaskReq", string(data)}, " ")
}

type CreateScheduleTaskReqScheduledType struct {
	value string
}

type CreateScheduleTaskReqScheduledTypeEnum struct {
	FIXED_TIME CreateScheduleTaskReqScheduledType
	DAY        CreateScheduleTaskReqScheduledType
	WEEK       CreateScheduleTaskReqScheduledType
	MONTH      CreateScheduleTaskReqScheduledType
}

func GetCreateScheduleTaskReqScheduledTypeEnum() CreateScheduleTaskReqScheduledTypeEnum {
	return CreateScheduleTaskReqScheduledTypeEnum{
		FIXED_TIME: CreateScheduleTaskReqScheduledType{
			value: "FIXED_TIME",
		},
		DAY: CreateScheduleTaskReqScheduledType{
			value: "DAY",
		},
		WEEK: CreateScheduleTaskReqScheduledType{
			value: "WEEK",
		},
		MONTH: CreateScheduleTaskReqScheduledType{
			value: "MONTH",
		},
	}
}

func (c CreateScheduleTaskReqScheduledType) Value() string {
	return c.value
}

func (c CreateScheduleTaskReqScheduledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScheduleTaskReqScheduledType) UnmarshalJSON(b []byte) error {
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
