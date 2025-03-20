package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LayerType 数仓分层的类型，可以是关系建模（ER）、维度建模（DIMENSION）和数据集市（DM）
type LayerType struct {
	value string
}

type LayerTypeEnum struct {
	THIRD_NF  LayerType
	DIMENSION LayerType
	DM        LayerType
}

func GetLayerTypeEnum() LayerTypeEnum {
	return LayerTypeEnum{
		THIRD_NF: LayerType{
			value: "THIRD_NF",
		},
		DIMENSION: LayerType{
			value: "DIMENSION",
		},
		DM: LayerType{
			value: "DM",
		},
	}
}

func (c LayerType) Value() string {
	return c.value
}

func (c LayerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LayerType) UnmarshalJSON(b []byte) error {
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
