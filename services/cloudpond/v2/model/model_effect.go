package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Effect 资源运营功能： DELETABLE：可删除 UNDELETABLE：不可删除
type Effect struct {
	value string
}

type EffectEnum struct {
	DELETABLE   Effect
	UNDELETABLE Effect
}

func GetEffectEnum() EffectEnum {
	return EffectEnum{
		DELETABLE: Effect{
			value: "DELETABLE",
		},
		UNDELETABLE: Effect{
			value: "UNDELETABLE",
		},
	}
}

func (c Effect) Value() string {
	return c.value
}

func (c Effect) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Effect) UnmarshalJSON(b []byte) error {
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
