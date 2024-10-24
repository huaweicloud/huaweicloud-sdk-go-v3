package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SpecScenario 规格使用场景，不填表示不限制： COMPUTE: 用于购买Ray计算资源时配置的物理节点规格 ENDPOINT: 用于创建Endpoint时配置的资源组规格大小
type SpecScenario struct {
	value string
}

type SpecScenarioEnum struct {
	COMPUTE  SpecScenario
	ENDPOINT SpecScenario
}

func GetSpecScenarioEnum() SpecScenarioEnum {
	return SpecScenarioEnum{
		COMPUTE: SpecScenario{
			value: "COMPUTE",
		},
		ENDPOINT: SpecScenario{
			value: "ENDPOINT",
		},
	}
}

func (c SpecScenario) Value() string {
	return c.value
}

func (c SpecScenario) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SpecScenario) UnmarshalJSON(b []byte) error {
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
