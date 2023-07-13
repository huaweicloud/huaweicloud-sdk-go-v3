package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourcesPlan struct {

	// 资源计划的周期类型，当前只允许以下类型：  daily
	PeriodType string `json:"period_type"`

	// 资源计划的起始时间，格式为“hour:minute”，表示时间在0:00-23:59之间。
	StartTime string `json:"start_time"`

	// 资源计划的结束时间，格式与“start_time”相同，不早于start_time表示的时间，且与start_time间隔不小于30min。
	EndTime string `json:"end_time"`

	// 资源计划内该节点组的最小保留节点数。 取值范围：[0～500]
	MinCapacity int32 `json:"min_capacity"`

	// 资源计划内该节点组的最大保留节点数。 取值范围：[0～500]
	MaxCapacity int32 `json:"max_capacity"`

	// 资源计划的生效日期，为空时代表每日，另外也可为以下返回值：  MONDAY（周一）、TUESDAY（周二）、WEDNESDAY（周三）、THURSDAY（周四）、FRIDAY（周五）、SATURDAY（周六）、SUNDAY（周日）
	EffectiveDays *[]ResourcesPlanEffectiveDays `json:"effective_days,omitempty"`
}

func (o ResourcesPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesPlan struct{}"
	}

	return strings.Join([]string{"ResourcesPlan", string(data)}, " ")
}

type ResourcesPlanEffectiveDays struct {
	value string
}

type ResourcesPlanEffectiveDaysEnum struct {
	MONDAY    ResourcesPlanEffectiveDays
	TUESDAY   ResourcesPlanEffectiveDays
	WEDNESDAY ResourcesPlanEffectiveDays
	THURSDAY  ResourcesPlanEffectiveDays
	FRIDAY    ResourcesPlanEffectiveDays
	SATURDAY  ResourcesPlanEffectiveDays
	SUNDAY    ResourcesPlanEffectiveDays
}

func GetResourcesPlanEffectiveDaysEnum() ResourcesPlanEffectiveDaysEnum {
	return ResourcesPlanEffectiveDaysEnum{
		MONDAY: ResourcesPlanEffectiveDays{
			value: "MONDAY",
		},
		TUESDAY: ResourcesPlanEffectiveDays{
			value: "TUESDAY",
		},
		WEDNESDAY: ResourcesPlanEffectiveDays{
			value: "WEDNESDAY",
		},
		THURSDAY: ResourcesPlanEffectiveDays{
			value: "THURSDAY",
		},
		FRIDAY: ResourcesPlanEffectiveDays{
			value: "FRIDAY",
		},
		SATURDAY: ResourcesPlanEffectiveDays{
			value: "SATURDAY",
		},
		SUNDAY: ResourcesPlanEffectiveDays{
			value: "SUNDAY",
		},
	}
}

func (c ResourcesPlanEffectiveDays) Value() string {
	return c.value
}

func (c ResourcesPlanEffectiveDays) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourcesPlanEffectiveDays) UnmarshalJSON(b []byte) error {
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
