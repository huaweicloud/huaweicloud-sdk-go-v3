package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateLifyCycleHookResponse Response Object
type CreateLifyCycleHookResponse struct {

	// 生命周期挂钩名称。
	LifecycleHookName *string `json:"lifecycle_hook_name,omitempty"`

	// 生命周期挂钩类型。INSTANCE_TERMINATING;INSTANCE_LAUNCHING
	LifecycleHookType *CreateLifyCycleHookResponseLifecycleHookType `json:"lifecycle_hook_type,omitempty"`

	// 生命周期挂钩默认回调操作。ABANDON;CONTINUE
	DefaultResult *CreateLifyCycleHookResponseDefaultResult `json:"default_result,omitempty"`

	// 生命周期挂钩超时时间，单位秒。
	DefaultTimeout *int32 `json:"default_timeout,omitempty"`

	// SMN服务中Topic的唯一的资源标识。
	NotificationTopicUrn *string `json:"notification_topic_urn,omitempty"`

	// SMN服务中Topic的资源名称。
	NotificationTopicName *string `json:"notification_topic_name,omitempty"`

	// 自定义通知消息。
	NotificationMetadata *string `json:"notification_metadata,omitempty"`

	// 生命周期挂钩创建时间，遵循UTC时间。
	CreateTime     *string `json:"create_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLifyCycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLifyCycleHookResponse struct{}"
	}

	return strings.Join([]string{"CreateLifyCycleHookResponse", string(data)}, " ")
}

type CreateLifyCycleHookResponseLifecycleHookType struct {
	value string
}

type CreateLifyCycleHookResponseLifecycleHookTypeEnum struct {
	INSTANCE_TERMINATING CreateLifyCycleHookResponseLifecycleHookType
	INSTANCE_LAUNCHING   CreateLifyCycleHookResponseLifecycleHookType
}

func GetCreateLifyCycleHookResponseLifecycleHookTypeEnum() CreateLifyCycleHookResponseLifecycleHookTypeEnum {
	return CreateLifyCycleHookResponseLifecycleHookTypeEnum{
		INSTANCE_TERMINATING: CreateLifyCycleHookResponseLifecycleHookType{
			value: "INSTANCE_TERMINATING",
		},
		INSTANCE_LAUNCHING: CreateLifyCycleHookResponseLifecycleHookType{
			value: "INSTANCE_LAUNCHING",
		},
	}
}

func (c CreateLifyCycleHookResponseLifecycleHookType) Value() string {
	return c.value
}

func (c CreateLifyCycleHookResponseLifecycleHookType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLifyCycleHookResponseLifecycleHookType) UnmarshalJSON(b []byte) error {
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

type CreateLifyCycleHookResponseDefaultResult struct {
	value string
}

type CreateLifyCycleHookResponseDefaultResultEnum struct {
	ABANDON  CreateLifyCycleHookResponseDefaultResult
	CONTINUE CreateLifyCycleHookResponseDefaultResult
}

func GetCreateLifyCycleHookResponseDefaultResultEnum() CreateLifyCycleHookResponseDefaultResultEnum {
	return CreateLifyCycleHookResponseDefaultResultEnum{
		ABANDON: CreateLifyCycleHookResponseDefaultResult{
			value: "ABANDON",
		},
		CONTINUE: CreateLifyCycleHookResponseDefaultResult{
			value: "CONTINUE",
		},
	}
}

func (c CreateLifyCycleHookResponseDefaultResult) Value() string {
	return c.value
}

func (c CreateLifyCycleHookResponseDefaultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLifyCycleHookResponseDefaultResult) UnmarshalJSON(b []byte) error {
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
