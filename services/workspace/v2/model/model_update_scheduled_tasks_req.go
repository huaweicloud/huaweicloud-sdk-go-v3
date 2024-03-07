package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateScheduledTasksReq 修改定时任务请求体。
type UpdateScheduledTasksReq struct {

	// 执行周期类型，可选值为： - FIXED_TIME：指定时间。 - DAY：按天。 - WEEK：按周。 - MONTH：按月。 - LIFE_CYCLE：指定场景下触发。
	ScheduledType *UpdateScheduledTasksReqScheduledType `json:"scheduled_type,omitempty"`

	// 周期按天时：按x天间隔执行。
	DayInterval *int32 `json:"day_interval,omitempty"`

	// 周期按周时：取值1~7，分别对应周日~周六，英文逗号分隔，如1,2,7。
	WeekList *string `json:"week_list,omitempty"`

	// 周期按月时：取值1~12，英文逗号分隔，如1,3,12。
	MonthList *string `json:"month_list,omitempty"`

	// 周期按月时：取值1~31及L(代表当月最后一天)，英文逗号分隔，如1,2,28,L。
	DateList *string `json:"date_list,omitempty"`

	// 周期指定时间时：表示指定的日期，格式为yyyy-MM-dd。
	ScheduledDate *string `json:"scheduled_date,omitempty"`

	// 指定的执行时间点，格式为HH:mm:ss。
	ScheduledTime *string `json:"scheduled_time,omitempty"`

	// 到期时间，格式为yyyy-MM-dd HH:mm:ss。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 灰度对象数量，优先级高于gray_desktop_ids。只支持脚本执行。
	GrayCount *int32 `json:"gray_count,omitempty"`

	// 灰度执行的桌面id列表，优先级低于gray_count。只支持脚本执行。
	GrayDesktopIds *[]string `json:"gray_desktop_ids,omitempty"`

	// 灰度失败阈值，灰度执行失败次数达到该值时，不执行下一批任务。只支持脚本执行。
	GrayFailThreshold *int32 `json:"gray_fail_threshold,omitempty"`

	// 触发场景。POST_CREATE_DESKTOP_SUCCESS：创建桌面成功后，POST_REBUILD_DESKTOP_SUCCESS：重建桌面成功后，POST_REATTACH_DESKTOP_SUCCESS：触发重建的分配用户任务成功后。
	LifeCycleType *string `json:"life_cycle_type,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 是否强制执行
	ForceExecute *bool `json:"force_execute,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 扩展参数，json格式。
	ExtraParams *string `json:"extra_params,omitempty"`

	// 定时任务应用的对象列表。
	ApplyObjects *[]ApplyObject `json:"apply_objects,omitempty"`

	// 优先级。触发式任务使用。
	Priority *int32 `json:"priority,omitempty"`

	// 任务是否灰度执行，供远程脚本使用。
	IsGray *bool `json:"is_gray,omitempty"`
}

func (o UpdateScheduledTasksReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledTasksReq struct{}"
	}

	return strings.Join([]string{"UpdateScheduledTasksReq", string(data)}, " ")
}

type UpdateScheduledTasksReqScheduledType struct {
	value string
}

type UpdateScheduledTasksReqScheduledTypeEnum struct {
	FIXED_TIME UpdateScheduledTasksReqScheduledType
	DAY        UpdateScheduledTasksReqScheduledType
	WEEK       UpdateScheduledTasksReqScheduledType
	MONTH      UpdateScheduledTasksReqScheduledType
	LIFE_CYCLE UpdateScheduledTasksReqScheduledType
}

func GetUpdateScheduledTasksReqScheduledTypeEnum() UpdateScheduledTasksReqScheduledTypeEnum {
	return UpdateScheduledTasksReqScheduledTypeEnum{
		FIXED_TIME: UpdateScheduledTasksReqScheduledType{
			value: "FIXED_TIME",
		},
		DAY: UpdateScheduledTasksReqScheduledType{
			value: "DAY",
		},
		WEEK: UpdateScheduledTasksReqScheduledType{
			value: "WEEK",
		},
		MONTH: UpdateScheduledTasksReqScheduledType{
			value: "MONTH",
		},
		LIFE_CYCLE: UpdateScheduledTasksReqScheduledType{
			value: "LIFE_CYCLE",
		},
	}
}

func (c UpdateScheduledTasksReqScheduledType) Value() string {
	return c.value
}

func (c UpdateScheduledTasksReqScheduledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScheduledTasksReqScheduledType) UnmarshalJSON(b []byte) error {
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
