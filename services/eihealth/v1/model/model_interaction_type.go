package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InteractionType 相互作用力类型：氢键、疏水作用、盐桥、π-π堆积、π-阳离子
type InteractionType struct {
	value string
}

type InteractionTypeEnum struct {
	HYDROGEN_BOND      InteractionType
	HYDROPHOBIC_ACTION InteractionType
	SALT_BRIDGE        InteractionType
	PI_STACKING        InteractionType
	PI_CATION          InteractionType
}

func GetInteractionTypeEnum() InteractionTypeEnum {
	return InteractionTypeEnum{
		HYDROGEN_BOND: InteractionType{
			value: "hydrogen_bond",
		},
		HYDROPHOBIC_ACTION: InteractionType{
			value: "hydrophobic_action",
		},
		SALT_BRIDGE: InteractionType{
			value: "salt_bridge",
		},
		PI_STACKING: InteractionType{
			value: "pi_stacking",
		},
		PI_CATION: InteractionType{
			value: "pi_cation",
		},
	}
}

func (c InteractionType) Value() string {
	return c.value
}

func (c InteractionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InteractionType) UnmarshalJSON(b []byte) error {
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
