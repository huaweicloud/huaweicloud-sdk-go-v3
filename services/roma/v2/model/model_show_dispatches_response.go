package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDispatchesResponse Response Object
type ShowDispatchesResponse struct {

	// 调度计划ID
	DispatchId *string `json:"dispatch_id,omitempty"`

	// 调度计划关联的任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 调度计划的执行开始时间
	StartDatetime *int64 `json:"start_datetime,omitempty"`

	// 调度计划执行周期的时间单位，当使用cron表达式时，为空 - MIN (分钟) - HOUR (小时) - DAY (日) - WEEK (周) - MON (月)
	Period *ShowDispatchesResponsePeriod `json:"period,omitempty"`

	// 调度计划的执行间隔时间周期
	DispatchInterval *int64 `json:"dispatch_interval,omitempty"`

	// 调度计划的创建时间
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 调度计划最近一次的修改时间
	LastModifiedDate *int64 `json:"last_modified_date,omitempty"`

	// 调度计划的备注信息
	Remark *string `json:"remark,omitempty"`

	// 调度计划是否使用cron表达式，允许如下值： - true (使用cron表达式) - false (不使用cron表达式)
	UseQuartzCron *bool `json:"use_quartz_cron,omitempty"`

	// 调度计划的cron表达式
	Cron           *string `json:"cron,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDispatchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDispatchesResponse struct{}"
	}

	return strings.Join([]string{"ShowDispatchesResponse", string(data)}, " ")
}

type ShowDispatchesResponsePeriod struct {
	value string
}

type ShowDispatchesResponsePeriodEnum struct {
	MIN  ShowDispatchesResponsePeriod
	HOUR ShowDispatchesResponsePeriod
	DAY  ShowDispatchesResponsePeriod
	WEEK ShowDispatchesResponsePeriod
	MON  ShowDispatchesResponsePeriod
}

func GetShowDispatchesResponsePeriodEnum() ShowDispatchesResponsePeriodEnum {
	return ShowDispatchesResponsePeriodEnum{
		MIN: ShowDispatchesResponsePeriod{
			value: "MIN",
		},
		HOUR: ShowDispatchesResponsePeriod{
			value: "HOUR",
		},
		DAY: ShowDispatchesResponsePeriod{
			value: "DAY",
		},
		WEEK: ShowDispatchesResponsePeriod{
			value: "WEEK",
		},
		MON: ShowDispatchesResponsePeriod{
			value: "MON",
		},
	}
}

func (c ShowDispatchesResponsePeriod) Value() string {
	return c.value
}

func (c ShowDispatchesResponsePeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDispatchesResponsePeriod) UnmarshalJSON(b []byte) error {
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
