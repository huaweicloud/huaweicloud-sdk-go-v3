package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UsageThreshold 用量阈值
type UsageThreshold struct {

	// 资源类型
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 原始资源类型
	SourceResourceSpecCode *string `json:"source_resource_spec_code,omitempty"`

	// 阈值
	Threshold float64 `json:"threshold"`

	// 阈值对应的单位,%,MB,GB 如果%，对应的阈值最大为95
	Unit UsageThresholdUnit `json:"unit"`

	// 开启或关闭当前资源类型的告警设置
	Enable *bool `json:"enable,omitempty"`
}

func (o UsageThreshold) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsageThreshold struct{}"
	}

	return strings.Join([]string{"UsageThreshold", string(data)}, " ")
}

type UsageThresholdUnit struct {
	value string
}

type UsageThresholdUnitEnum struct {
	PERCENT UsageThresholdUnit
	MB      UsageThresholdUnit
	GB      UsageThresholdUnit
}

func GetUsageThresholdUnitEnum() UsageThresholdUnitEnum {
	return UsageThresholdUnitEnum{
		PERCENT: UsageThresholdUnit{
			value: "%",
		},
		MB: UsageThresholdUnit{
			value: "MB",
		},
		GB: UsageThresholdUnit{
			value: "GB",
		},
	}
}

func (c UsageThresholdUnit) Value() string {
	return c.value
}

func (c UsageThresholdUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UsageThresholdUnit) UnmarshalJSON(b []byte) error {
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
