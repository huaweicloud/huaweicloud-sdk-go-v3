package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiActionDto struct {

	// 截止时间
	Time *string `json:"time,omitempty"`

	// 操作类型
	Action *ApiActionDtoAction `json:"action,omitempty"`
}

func (o ApiActionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiActionDto struct{}"
	}

	return strings.Join([]string{"ApiActionDto", string(data)}, " ")
}

type ApiActionDtoAction struct {
	value string
}

type ApiActionDtoActionEnum struct {
	UNPUBLISH ApiActionDtoAction
	STOP      ApiActionDtoAction
	RECOVER   ApiActionDtoAction
}

func GetApiActionDtoActionEnum() ApiActionDtoActionEnum {
	return ApiActionDtoActionEnum{
		UNPUBLISH: ApiActionDtoAction{
			value: "UNPUBLISH",
		},
		STOP: ApiActionDtoAction{
			value: "STOP",
		},
		RECOVER: ApiActionDtoAction{
			value: "RECOVER",
		},
	}
}

func (c ApiActionDtoAction) Value() string {
	return c.value
}

func (c ApiActionDtoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiActionDtoAction) UnmarshalJSON(b []byte) error {
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
