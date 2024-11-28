package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateDesktopPoolAuthorizedObjectsRequestBody struct {

	// 要授权的用户/用户组。
	Objects *[]AuthorizedObjects `json:"objects,omitempty"`

	// 执行动作，ADD：增加授权用户/用户组，REMOVE：移除已授权用户/用户组
	Action CreateDesktopPoolAuthorizedObjectsRequestBodyAction `json:"action"`
}

func (o CreateDesktopPoolAuthorizedObjectsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopPoolAuthorizedObjectsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDesktopPoolAuthorizedObjectsRequestBody", string(data)}, " ")
}

type CreateDesktopPoolAuthorizedObjectsRequestBodyAction struct {
	value string
}

type CreateDesktopPoolAuthorizedObjectsRequestBodyActionEnum struct {
	ADD    CreateDesktopPoolAuthorizedObjectsRequestBodyAction
	REMOVE CreateDesktopPoolAuthorizedObjectsRequestBodyAction
}

func GetCreateDesktopPoolAuthorizedObjectsRequestBodyActionEnum() CreateDesktopPoolAuthorizedObjectsRequestBodyActionEnum {
	return CreateDesktopPoolAuthorizedObjectsRequestBodyActionEnum{
		ADD: CreateDesktopPoolAuthorizedObjectsRequestBodyAction{
			value: "ADD",
		},
		REMOVE: CreateDesktopPoolAuthorizedObjectsRequestBodyAction{
			value: "REMOVE",
		},
	}
}

func (c CreateDesktopPoolAuthorizedObjectsRequestBodyAction) Value() string {
	return c.value
}

func (c CreateDesktopPoolAuthorizedObjectsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDesktopPoolAuthorizedObjectsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
