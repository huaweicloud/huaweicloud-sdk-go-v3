package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PromLimits struct {

	// 指标存储时长，只支持 15天，30天，60天 ，90天
	CompactorBlocksRetentionPeriod PromLimitsCompactorBlocksRetentionPeriod `json:"compactor_blocks_retention_period"`
}

func (o PromLimits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PromLimits struct{}"
	}

	return strings.Join([]string{"PromLimits", string(data)}, " ")
}

type PromLimitsCompactorBlocksRetentionPeriod struct {
	value string
}

type PromLimitsCompactorBlocksRetentionPeriodEnum struct {
	E_360H  PromLimitsCompactorBlocksRetentionPeriod
	E_720H  PromLimitsCompactorBlocksRetentionPeriod
	E_1440H PromLimitsCompactorBlocksRetentionPeriod
	E_2160H PromLimitsCompactorBlocksRetentionPeriod
}

func GetPromLimitsCompactorBlocksRetentionPeriodEnum() PromLimitsCompactorBlocksRetentionPeriodEnum {
	return PromLimitsCompactorBlocksRetentionPeriodEnum{
		E_360H: PromLimitsCompactorBlocksRetentionPeriod{
			value: "\"360h\"",
		},
		E_720H: PromLimitsCompactorBlocksRetentionPeriod{
			value: " \"720h\"",
		},
		E_1440H: PromLimitsCompactorBlocksRetentionPeriod{
			value: " \"1440h\"",
		},
		E_2160H: PromLimitsCompactorBlocksRetentionPeriod{
			value: " \"2160h\"",
		},
	}
}

func (c PromLimitsCompactorBlocksRetentionPeriod) Value() string {
	return c.value
}

func (c PromLimitsCompactorBlocksRetentionPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PromLimitsCompactorBlocksRetentionPeriod) UnmarshalJSON(b []byte) error {
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
