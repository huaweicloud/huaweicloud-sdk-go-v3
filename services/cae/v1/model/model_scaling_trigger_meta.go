package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScalingTriggerMeta trigger元数据。
type ScalingTriggerMeta struct {

	// 数据类型，当前只支持利用率，默认值为Utilization。  type为\"cpu、memory\"时，配置此参数。
	Type *string `json:"type,omitempty"`

	// 触发指标的阈值。  type为\"cpu、memory\"时，配置此参数。
	Value *string `json:"value,omitempty"`

	// 生效周期。  type为\"cron\"时，配置此参数。
	PeriodType *ScalingTriggerMetaPeriodType `json:"period_type,omitempty"`

	// 每个周期内触发的时间点和实例数。  type为\"cron\"时，配置此参数。
	Schedulers *[]CronTriggerScheduler `json:"schedulers,omitempty"`
}

func (o ScalingTriggerMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingTriggerMeta struct{}"
	}

	return strings.Join([]string{"ScalingTriggerMeta", string(data)}, " ")
}

type ScalingTriggerMetaPeriodType struct {
	value string
}

type ScalingTriggerMetaPeriodTypeEnum struct {
	DAY   ScalingTriggerMetaPeriodType
	WEEK  ScalingTriggerMetaPeriodType
	MONTH ScalingTriggerMetaPeriodType
}

func GetScalingTriggerMetaPeriodTypeEnum() ScalingTriggerMetaPeriodTypeEnum {
	return ScalingTriggerMetaPeriodTypeEnum{
		DAY: ScalingTriggerMetaPeriodType{
			value: "day",
		},
		WEEK: ScalingTriggerMetaPeriodType{
			value: "week",
		},
		MONTH: ScalingTriggerMetaPeriodType{
			value: "month",
		},
	}
}

func (c ScalingTriggerMetaPeriodType) Value() string {
	return c.value
}

func (c ScalingTriggerMetaPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingTriggerMetaPeriodType) UnmarshalJSON(b []byte) error {
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
