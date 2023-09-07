package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ActionKindObj API类型，固定值“Action”，该值不可修改。
type ActionKindObj struct {
	value string
}

type ActionKindObjEnum struct {
	ACTION ActionKindObj
}

func GetActionKindObjEnum() ActionKindObjEnum {
	return ActionKindObjEnum{
		ACTION: ActionKindObj{
			value: "Action",
		},
	}
}

func (c ActionKindObj) Value() string {
	return c.value
}

func (c ActionKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionKindObj) UnmarshalJSON(b []byte) error {
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
