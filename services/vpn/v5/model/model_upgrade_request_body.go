package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpgradeRequestBody struct {

	// 升级操作，可选值start、finish、rollback
	Action UpgradeRequestBodyAction `json:"action"`
}

func (o UpgradeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeRequestBody", string(data)}, " ")
}

type UpgradeRequestBodyAction struct {
	value string
}

type UpgradeRequestBodyActionEnum struct {
	START    UpgradeRequestBodyAction
	FINISH   UpgradeRequestBodyAction
	ROLLBACK UpgradeRequestBodyAction
}

func GetUpgradeRequestBodyActionEnum() UpgradeRequestBodyActionEnum {
	return UpgradeRequestBodyActionEnum{
		START: UpgradeRequestBodyAction{
			value: "start",
		},
		FINISH: UpgradeRequestBodyAction{
			value: "finish",
		},
		ROLLBACK: UpgradeRequestBodyAction{
			value: "rollback",
		},
	}
}

func (c UpgradeRequestBodyAction) Value() string {
	return c.value
}

func (c UpgradeRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradeRequestBodyAction) UnmarshalJSON(b []byte) error {
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
