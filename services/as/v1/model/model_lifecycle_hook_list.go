package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LifecycleHookList 生命周期挂钩
type LifecycleHookList struct {

	// 生命周期挂钩名称。
	LifecycleHookName *string `json:"lifecycle_hook_name,omitempty"`

	// 生命周期挂钩类型。INSTANCE_TERMINATING；INSTANCE_LAUNCHING。
	LifecycleHookType *LifecycleHookListLifecycleHookType `json:"lifecycle_hook_type,omitempty"`

	// 生命周期挂钩默认回调操作。ABANDON;CONTINUE。
	DefaultResult *LifecycleHookListDefaultResult `json:"default_result,omitempty"`

	// 生命周期挂钩超时时间，单位秒。
	DefaultTimeout *int32 `json:"default_timeout,omitempty"`

	// SMN服务中Topic的唯一的资源标识。
	NotificationTopicUrn *string `json:"notification_topic_urn,omitempty"`

	// SMN服务中Topic的资源名称。
	NotificationTopicName *string `json:"notification_topic_name,omitempty"`

	// 自定义通知消息。
	NotificationMetadata *string `json:"notification_metadata,omitempty"`

	// 创建生命周期挂钩时间，遵循UTC时间。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o LifecycleHookList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleHookList struct{}"
	}

	return strings.Join([]string{"LifecycleHookList", string(data)}, " ")
}

type LifecycleHookListLifecycleHookType struct {
	value string
}

type LifecycleHookListLifecycleHookTypeEnum struct {
	INSTANCE_TERMINATING LifecycleHookListLifecycleHookType
	INSTANCE_LAUNCHING   LifecycleHookListLifecycleHookType
}

func GetLifecycleHookListLifecycleHookTypeEnum() LifecycleHookListLifecycleHookTypeEnum {
	return LifecycleHookListLifecycleHookTypeEnum{
		INSTANCE_TERMINATING: LifecycleHookListLifecycleHookType{
			value: "INSTANCE_TERMINATING",
		},
		INSTANCE_LAUNCHING: LifecycleHookListLifecycleHookType{
			value: "INSTANCE_LAUNCHING",
		},
	}
}

func (c LifecycleHookListLifecycleHookType) Value() string {
	return c.value
}

func (c LifecycleHookListLifecycleHookType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LifecycleHookListLifecycleHookType) UnmarshalJSON(b []byte) error {
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

type LifecycleHookListDefaultResult struct {
	value string
}

type LifecycleHookListDefaultResultEnum struct {
	ABANDON  LifecycleHookListDefaultResult
	CONTINUE LifecycleHookListDefaultResult
}

func GetLifecycleHookListDefaultResultEnum() LifecycleHookListDefaultResultEnum {
	return LifecycleHookListDefaultResultEnum{
		ABANDON: LifecycleHookListDefaultResult{
			value: "ABANDON",
		},
		CONTINUE: LifecycleHookListDefaultResult{
			value: "CONTINUE",
		},
	}
}

func (c LifecycleHookListDefaultResult) Value() string {
	return c.value
}

func (c LifecycleHookListDefaultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LifecycleHookListDefaultResult) UnmarshalJSON(b []byte) error {
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
