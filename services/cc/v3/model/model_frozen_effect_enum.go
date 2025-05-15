package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FrozenEffectEnum 冻结效果。 - RELEASABLE（冻结可释放） - UNRELEASABLE（冻结不可释放）
type FrozenEffectEnum struct {
	value string
}

type FrozenEffectEnumEnum struct {
	RELEASABLE   FrozenEffectEnum
	UNRELEASABLE FrozenEffectEnum
}

func GetFrozenEffectEnumEnum() FrozenEffectEnumEnum {
	return FrozenEffectEnumEnum{
		RELEASABLE: FrozenEffectEnum{
			value: "RELEASABLE",
		},
		UNRELEASABLE: FrozenEffectEnum{
			value: "UNRELEASABLE",
		},
	}
}

func (c FrozenEffectEnum) Value() string {
	return c.value
}

func (c FrozenEffectEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FrozenEffectEnum) UnmarshalJSON(b []byte) error {
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
