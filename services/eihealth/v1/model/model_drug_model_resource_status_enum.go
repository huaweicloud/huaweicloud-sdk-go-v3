package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DrugModelResourceStatusEnum **参数解释**： 模型状态。 **约束限制**： 不涉及 **取值范围**： - Normal：正常。 - Freeze：冻结。 - Deleted：删除。 **默认取值**： 不涉及
type DrugModelResourceStatusEnum struct {
	value string
}

type DrugModelResourceStatusEnumEnum struct {
	NORMAL  DrugModelResourceStatusEnum
	FREEZE  DrugModelResourceStatusEnum
	DELETED DrugModelResourceStatusEnum
}

func GetDrugModelResourceStatusEnumEnum() DrugModelResourceStatusEnumEnum {
	return DrugModelResourceStatusEnumEnum{
		NORMAL: DrugModelResourceStatusEnum{
			value: "Normal",
		},
		FREEZE: DrugModelResourceStatusEnum{
			value: "Freeze",
		},
		DELETED: DrugModelResourceStatusEnum{
			value: "Deleted",
		},
	}
}

func (c DrugModelResourceStatusEnum) Value() string {
	return c.value
}

func (c DrugModelResourceStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DrugModelResourceStatusEnum) UnmarshalJSON(b []byte) error {
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
