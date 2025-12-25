package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HealthStatus **参数解释**: 采集通道监控状态 - SUCCESS_PART 部分运行 - SUCCESS_ALL 全部运行 - SUCCESS_NO 没有运行  **约束限制** 不涉及 **取值范围**: - SUCCESS_PART - SUCCESS_ALL - SUCCESS_NO  **默认值** 不涉及
type HealthStatus struct {
	value string
}

type HealthStatusEnum struct {
	SUCCESS_PART HealthStatus
	SUCCESS_ALL  HealthStatus
	SUCCESS_NO   HealthStatus
}

func GetHealthStatusEnum() HealthStatusEnum {
	return HealthStatusEnum{
		SUCCESS_PART: HealthStatus{
			value: "SUCCESS_PART",
		},
		SUCCESS_ALL: HealthStatus{
			value: "SUCCESS_ALL",
		},
		SUCCESS_NO: HealthStatus{
			value: "SUCCESS_NO",
		},
	}
}

func (c HealthStatus) Value() string {
	return c.value
}

func (c HealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HealthStatus) UnmarshalJSON(b []byte) error {
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
