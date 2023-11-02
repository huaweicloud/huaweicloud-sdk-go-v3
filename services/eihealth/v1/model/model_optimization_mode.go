package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OptimizationMode 侧链修饰
type OptimizationMode struct {
	value string
}

type OptimizationModeEnum struct {
	GENERATION             OptimizationMode
	SIDE_CHAINS_DECORATION OptimizationMode
}

func GetOptimizationModeEnum() OptimizationModeEnum {
	return OptimizationModeEnum{
		GENERATION: OptimizationMode{
			value: "generation",
		},
		SIDE_CHAINS_DECORATION: OptimizationMode{
			value: "side_chains_decoration",
		},
	}
}

func (c OptimizationMode) Value() string {
	return c.value
}

func (c OptimizationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OptimizationMode) UnmarshalJSON(b []byte) error {
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
