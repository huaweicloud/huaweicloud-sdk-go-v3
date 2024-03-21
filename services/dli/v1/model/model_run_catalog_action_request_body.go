package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RunCatalogActionRequestBody struct {

	// catalog操作:bind或者unbind。
	Action RunCatalogActionRequestBodyAction `json:"action"`

	// DLI侧catalog映射名称.
	Name string `json:"name"`

	Parameters map[string]string `json:"parameters"`
}

func (o RunCatalogActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCatalogActionRequestBody struct{}"
	}

	return strings.Join([]string{"RunCatalogActionRequestBody", string(data)}, " ")
}

type RunCatalogActionRequestBodyAction struct {
	value string
}

type RunCatalogActionRequestBodyActionEnum struct {
	BIND   RunCatalogActionRequestBodyAction
	UNBIND RunCatalogActionRequestBodyAction
}

func GetRunCatalogActionRequestBodyActionEnum() RunCatalogActionRequestBodyActionEnum {
	return RunCatalogActionRequestBodyActionEnum{
		BIND: RunCatalogActionRequestBodyAction{
			value: "bind",
		},
		UNBIND: RunCatalogActionRequestBodyAction{
			value: "unbind",
		},
	}
}

func (c RunCatalogActionRequestBodyAction) Value() string {
	return c.value
}

func (c RunCatalogActionRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunCatalogActionRequestBodyAction) UnmarshalJSON(b []byte) error {
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
