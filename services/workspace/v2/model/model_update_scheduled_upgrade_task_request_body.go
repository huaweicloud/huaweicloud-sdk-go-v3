package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateScheduledUpgradeTaskRequestBody 更新升级任务请求
type UpdateScheduledUpgradeTaskRequestBody struct {

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 执行周期类型：FIXED_TIME-指定时间 DAY-按天 WEEK-按周 MONTH-按月
	ScheduledType *UpdateScheduledUpgradeTaskRequestBodyScheduledType `json:"scheduled_type,omitempty"`

	// 时区
	Timezone *string `json:"timezone,omitempty"`

	// 周期按周时：取值1~7，英文逗号分隔，如1,2,7
	WeekList *string `json:"week_list,omitempty"`

	// 周期按月时：取值1~12，英文逗号分隔
	MonthList *string `json:"month_list,omitempty"`

	// 周期按月时：取值1~31及L(代表当月最后一天)
	DateList *string `json:"date_list,omitempty"`

	// 按天跳过天数
	DayInterval *int32 `json:"day_interval,omitempty"`

	// 周期指定时间时：表示指定的日期
	ScheduledDate *string `json:"scheduled_date,omitempty"`

	// 指定的执行时间点
	ScheduledTime *string `json:"scheduled_time,omitempty"`

	// 是否强制升级：0-否 1-是
	IsForceExecute *int32 `json:"is_force_execute,omitempty"`

	// 低于此版本升级
	MinVersion *string `json:"min_version,omitempty"`

	// 升级目标版本
	TargetVersion *string `json:"target_version,omitempty"`

	// 过期时间开启：0-未开启 1-开启
	ExpireEnable *int32 `json:"expire_enable,omitempty"`

	// 过期时间
	ExpireTime *string `json:"expire_time,omitempty"`

	// 是否通知：0-不通知 1-通知
	IsNotify *int32 `json:"is_notify,omitempty"`

	// 扩展参数（JSON格式）
	ExtraParams *string `json:"extra_params,omitempty"`

	// 执行策略：0-全量下发 1-灰度下发
	ExecuteStrategy *int32 `json:"execute_strategy,omitempty"`

	// 灰度规则：0-确定 1-随机（execute_strategy=1时使用）
	GrayscaleRule *int32 `json:"grayscale_rule,omitempty"`

	// 随机首批执行数
	RandomFirstBatchCount *int32 `json:"random_first_batch_count,omitempty"`

	// 灰度对象id列表（JSON数组格式）
	GrayObjectIds *string `json:"gray_object_ids,omitempty"`

	// 首批执行失败阈值
	GrayFailThreshold *int32 `json:"gray_fail_threshold,omitempty"`

	// 时间窗结束时间
	ScheduledEndTime *string `json:"scheduled_end_time,omitempty"`

	// 任务描述
	Description *string `json:"description,omitempty"`

	// 启用状态：0-未启用 1-启用
	IsEnable *int32 `json:"is_enable,omitempty"`

	// 应用对象列表
	ApplyObjects *[]TaskApplyObjectInfo `json:"apply_objects,omitempty"`
}

func (o UpdateScheduledUpgradeTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledUpgradeTaskRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateScheduledUpgradeTaskRequestBody", string(data)}, " ")
}

type UpdateScheduledUpgradeTaskRequestBodyScheduledType struct {
	value string
}

type UpdateScheduledUpgradeTaskRequestBodyScheduledTypeEnum struct {
	FIXED_TIME UpdateScheduledUpgradeTaskRequestBodyScheduledType
	DAY        UpdateScheduledUpgradeTaskRequestBodyScheduledType
	WEEK       UpdateScheduledUpgradeTaskRequestBodyScheduledType
	MONTH      UpdateScheduledUpgradeTaskRequestBodyScheduledType
}

func GetUpdateScheduledUpgradeTaskRequestBodyScheduledTypeEnum() UpdateScheduledUpgradeTaskRequestBodyScheduledTypeEnum {
	return UpdateScheduledUpgradeTaskRequestBodyScheduledTypeEnum{
		FIXED_TIME: UpdateScheduledUpgradeTaskRequestBodyScheduledType{
			value: "FIXED_TIME",
		},
		DAY: UpdateScheduledUpgradeTaskRequestBodyScheduledType{
			value: "DAY",
		},
		WEEK: UpdateScheduledUpgradeTaskRequestBodyScheduledType{
			value: "WEEK",
		},
		MONTH: UpdateScheduledUpgradeTaskRequestBodyScheduledType{
			value: "MONTH",
		},
	}
}

func (c UpdateScheduledUpgradeTaskRequestBodyScheduledType) Value() string {
	return c.value
}

func (c UpdateScheduledUpgradeTaskRequestBodyScheduledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScheduledUpgradeTaskRequestBodyScheduledType) UnmarshalJSON(b []byte) error {
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
