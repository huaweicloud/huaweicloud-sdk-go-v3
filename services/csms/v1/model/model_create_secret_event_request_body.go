package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSecretEventRequestBody 创建凭据事件通知请求体。
type CreateSecretEventRequestBody struct {

	// 新创建事件通知的名称。  约束：取值范围为1到64个字符，满足正则匹配“^[a-zA-Z0-9_-]{1,64}$”。
	Name string `json:"name"`

	// 本次事件通知的基础事件列表，基础事件类型如下。  - SECRET_VERSION_CREATED：版本创建 - SECRET_VERSION_EXPIRED：版本过期 - SECRET_ROTATED：凭据轮转 - SECRET_DELETED：凭据删除 - SECRET_ROTATED_FAILED:凭据轮转失败  列表包含的基础事件类型不能重复。
	EventTypes []CreateSecretEventRequestBodyEventTypes `json:"event_types"`

	// 控制事件是否生效，只有启用状态才能触发包含的基础事件类型  ENABLED：启用 DISABLED：禁用
	State CreateSecretEventRequestBodyState `json:"state"`

	Notification *Notification `json:"notification"`
}

func (o CreateSecretEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretEventRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecretEventRequestBody", string(data)}, " ")
}

type CreateSecretEventRequestBodyEventTypes struct {
	value string
}

type CreateSecretEventRequestBodyEventTypesEnum struct {
	SECRET_VERSION_CREATED CreateSecretEventRequestBodyEventTypes
	SECRET_VERSION_EXPIRED CreateSecretEventRequestBodyEventTypes
	SECRET_ROTATED         CreateSecretEventRequestBodyEventTypes
	SECRET_DELETED         CreateSecretEventRequestBodyEventTypes
	SECRET_ROTATED_FAILED  CreateSecretEventRequestBodyEventTypes
}

func GetCreateSecretEventRequestBodyEventTypesEnum() CreateSecretEventRequestBodyEventTypesEnum {
	return CreateSecretEventRequestBodyEventTypesEnum{
		SECRET_VERSION_CREATED: CreateSecretEventRequestBodyEventTypes{
			value: "SECRET_VERSION_CREATED",
		},
		SECRET_VERSION_EXPIRED: CreateSecretEventRequestBodyEventTypes{
			value: "SECRET_VERSION_EXPIRED",
		},
		SECRET_ROTATED: CreateSecretEventRequestBodyEventTypes{
			value: "SECRET_ROTATED",
		},
		SECRET_DELETED: CreateSecretEventRequestBodyEventTypes{
			value: "SECRET_DELETED",
		},
		SECRET_ROTATED_FAILED: CreateSecretEventRequestBodyEventTypes{
			value: "SECRET_ROTATED_FAILED",
		},
	}
}

func (c CreateSecretEventRequestBodyEventTypes) Value() string {
	return c.value
}

func (c CreateSecretEventRequestBodyEventTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecretEventRequestBodyEventTypes) UnmarshalJSON(b []byte) error {
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

type CreateSecretEventRequestBodyState struct {
	value string
}

type CreateSecretEventRequestBodyStateEnum struct {
	ENABLED  CreateSecretEventRequestBodyState
	DISABLED CreateSecretEventRequestBodyState
}

func GetCreateSecretEventRequestBodyStateEnum() CreateSecretEventRequestBodyStateEnum {
	return CreateSecretEventRequestBodyStateEnum{
		ENABLED: CreateSecretEventRequestBodyState{
			value: "ENABLED",
		},
		DISABLED: CreateSecretEventRequestBodyState{
			value: "DISABLED",
		},
	}
}

func (c CreateSecretEventRequestBodyState) Value() string {
	return c.value
}

func (c CreateSecretEventRequestBodyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecretEventRequestBodyState) UnmarshalJSON(b []byte) error {
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
