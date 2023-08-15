package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Topic struct {

	// TOPIC的ID
	Id *int32 `json:"id,omitempty"`

	// TOPIC的名称
	Name *string `json:"name,omitempty"`

	// TOPIC描述
	Description *string `json:"description,omitempty"`

	// TOPIC权限, 主题权限 0-发布 1-订阅
	Permission *TopicPermission `json:"permission,omitempty"`

	// TOPIC类型 0-基础TOPIC 1-用户自定义TOPIC
	IsPrivate *TopicIsPrivate `json:"is_private,omitempty"`
}

func (o Topic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Topic struct{}"
	}

	return strings.Join([]string{"Topic", string(data)}, " ")
}

type TopicPermission struct {
	value int32
}

type TopicPermissionEnum struct {
	E_0 TopicPermission
	E_1 TopicPermission
}

func GetTopicPermissionEnum() TopicPermissionEnum {
	return TopicPermissionEnum{
		E_0: TopicPermission{
			value: 0,
		}, E_1: TopicPermission{
			value: 1,
		},
	}
}

func (c TopicPermission) Value() int32 {
	return c.value
}

func (c TopicPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TopicPermission) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type TopicIsPrivate struct {
	value int32
}

type TopicIsPrivateEnum struct {
	E_0 TopicIsPrivate
	E_1 TopicIsPrivate
}

func GetTopicIsPrivateEnum() TopicIsPrivateEnum {
	return TopicIsPrivateEnum{
		E_0: TopicIsPrivate{
			value: 0,
		}, E_1: TopicIsPrivate{
			value: 1,
		},
	}
}

func (c TopicIsPrivate) Value() int32 {
	return c.value
}

func (c TopicIsPrivate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TopicIsPrivate) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
