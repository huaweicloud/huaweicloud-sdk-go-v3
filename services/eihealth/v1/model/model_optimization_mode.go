package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OptimizationMode 靶点口袋分子设计模式：支持从头生成、侧链修饰、骨架跃迁、片段生长。
type OptimizationMode struct {
	value string
}

type OptimizationModeEnum struct {
	GENERATION             OptimizationMode
	SIDE_CHAINS_DECORATION OptimizationMode
	SCAFFOLD_HOPPING       OptimizationMode
	FRAGMENT_GROWING       OptimizationMode
}

func GetOptimizationModeEnum() OptimizationModeEnum {
	return OptimizationModeEnum{
		GENERATION: OptimizationMode{
			value: "generation",
		},
		SIDE_CHAINS_DECORATION: OptimizationMode{
			value: "side_chains_decoration",
		},
		SCAFFOLD_HOPPING: OptimizationMode{
			value: "scaffold_hopping",
		},
		FRAGMENT_GROWING: OptimizationMode{
			value: "fragment_growing",
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
