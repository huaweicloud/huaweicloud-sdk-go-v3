package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSecretEventRequestBody 更新凭据事件通知请求体。
type UpdateSecretEventRequestBody struct {

	// 事件通知状态，取值如下。  ENABLED：表示启用状态 DISABLED：表示禁用状态
	State *UpdateSecretEventRequestBodyState `json:"state,omitempty"`

	// 本次事件通知的基础事件列表，基础事件类型如下。  - SECRET_VERSION_CREATED:版本创建 - SECRET_VERSION_EXPIRED:版本过期 - SECRET_ROTATED:凭据轮转成功 - SECRET_DELETED:凭据删除 - SECRET_ROTATED_FAILED:凭据轮转失败  列表包含的基础事件类型不能重复。
	EventTypes *[]UpdateSecretEventRequestBodyEventTypes `json:"event_types,omitempty"`

	Notification *Notification `json:"notification,omitempty"`
}

func (o UpdateSecretEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretEventRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecretEventRequestBody", string(data)}, " ")
}

type UpdateSecretEventRequestBodyState struct {
	value string
}

type UpdateSecretEventRequestBodyStateEnum struct {
	ENABLED  UpdateSecretEventRequestBodyState
	DISABLED UpdateSecretEventRequestBodyState
}

func GetUpdateSecretEventRequestBodyStateEnum() UpdateSecretEventRequestBodyStateEnum {
	return UpdateSecretEventRequestBodyStateEnum{
		ENABLED: UpdateSecretEventRequestBodyState{
			value: "ENABLED",
		},
		DISABLED: UpdateSecretEventRequestBodyState{
			value: "DISABLED",
		},
	}
}

func (c UpdateSecretEventRequestBodyState) Value() string {
	return c.value
}

func (c UpdateSecretEventRequestBodyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecretEventRequestBodyState) UnmarshalJSON(b []byte) error {
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

type UpdateSecretEventRequestBodyEventTypes struct {
	value string
}

type UpdateSecretEventRequestBodyEventTypesEnum struct {
	SECRET_VERSION_CREATED UpdateSecretEventRequestBodyEventTypes
	SECRET_VERSION_EXPIRED UpdateSecretEventRequestBodyEventTypes
	SECRET_ROTATED         UpdateSecretEventRequestBodyEventTypes
	SECRET_DELETED         UpdateSecretEventRequestBodyEventTypes
	SECRET_ROTATED_FAILED  UpdateSecretEventRequestBodyEventTypes
}

func GetUpdateSecretEventRequestBodyEventTypesEnum() UpdateSecretEventRequestBodyEventTypesEnum {
	return UpdateSecretEventRequestBodyEventTypesEnum{
		SECRET_VERSION_CREATED: UpdateSecretEventRequestBodyEventTypes{
			value: "SECRET_VERSION_CREATED",
		},
		SECRET_VERSION_EXPIRED: UpdateSecretEventRequestBodyEventTypes{
			value: "SECRET_VERSION_EXPIRED",
		},
		SECRET_ROTATED: UpdateSecretEventRequestBodyEventTypes{
			value: "SECRET_ROTATED",
		},
		SECRET_DELETED: UpdateSecretEventRequestBodyEventTypes{
			value: "SECRET_DELETED",
		},
		SECRET_ROTATED_FAILED: UpdateSecretEventRequestBodyEventTypes{
			value: "SECRET_ROTATED_FAILED",
		},
	}
}

func (c UpdateSecretEventRequestBodyEventTypes) Value() string {
	return c.value
}

func (c UpdateSecretEventRequestBodyEventTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecretEventRequestBodyEventTypes) UnmarshalJSON(b []byte) error {
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
