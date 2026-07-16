package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAuthModeRequestBody struct {

	// **参数解释**：更新的模式类型。 **取值范围**： - strict：严格模式。 - loose：非严格模式。
	Mode UpdateAuthModeRequestBodyMode `json:"mode"`
}

func (o UpdateAuthModeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthModeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAuthModeRequestBody", string(data)}, " ")
}

type UpdateAuthModeRequestBodyMode struct {
	value string
}

type UpdateAuthModeRequestBodyModeEnum struct {
	STRICT UpdateAuthModeRequestBodyMode
	LOOSE  UpdateAuthModeRequestBodyMode
}

func GetUpdateAuthModeRequestBodyModeEnum() UpdateAuthModeRequestBodyModeEnum {
	return UpdateAuthModeRequestBodyModeEnum{
		STRICT: UpdateAuthModeRequestBodyMode{
			value: "strict",
		},
		LOOSE: UpdateAuthModeRequestBodyMode{
			value: "loose",
		},
	}
}

func (c UpdateAuthModeRequestBodyMode) Value() string {
	return c.value
}

func (c UpdateAuthModeRequestBodyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAuthModeRequestBodyMode) UnmarshalJSON(b []byte) error {
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
