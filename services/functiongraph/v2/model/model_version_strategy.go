package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VersionStrategy 灰度策略信息
type VersionStrategy struct {

	// 规则列表
	Rules *[]VersionStrategyRules `json:"rules,omitempty"`

	// 所有规则聚合方式。and：所有规则都满足，or：满足其中一个
	CombineType *VersionStrategyCombineType `json:"combine_type,omitempty"`
}

func (o VersionStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionStrategy struct{}"
	}

	return strings.Join([]string{"VersionStrategy", string(data)}, " ")
}

type VersionStrategyCombineType struct {
	value string
}

type VersionStrategyCombineTypeEnum struct {
	AND VersionStrategyCombineType
	OR  VersionStrategyCombineType
}

func GetVersionStrategyCombineTypeEnum() VersionStrategyCombineTypeEnum {
	return VersionStrategyCombineTypeEnum{
		AND: VersionStrategyCombineType{
			value: "and",
		},
		OR: VersionStrategyCombineType{
			value: "or",
		},
	}
}

func (c VersionStrategyCombineType) Value() string {
	return c.value
}

func (c VersionStrategyCombineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionStrategyCombineType) UnmarshalJSON(b []byte) error {
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
