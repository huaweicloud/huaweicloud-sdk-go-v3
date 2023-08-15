package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTopicReq struct {

	// 主题名称，只能由英文字母、数字、百分号、竖线、中划线、下划线组成，长度3~64个字符。
	Name *string `json:"name,omitempty"`

	// 关联的代理。
	Brokers *[]string `json:"brokers,omitempty"`

	// 队列数，范围1~50。
	QueueNum float32 `json:"queue_num,omitempty"`

	// 权限。
	Permission *CreateTopicReqPermission `json:"permission,omitempty"`

	// 消息类型。
	MessageType *CreateTopicReqMessageType `json:"message_type,omitempty"`
}

func (o CreateTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicReq struct{}"
	}

	return strings.Join([]string{"CreateTopicReq", string(data)}, " ")
}

type CreateTopicReqPermission struct {
	value string
}

type CreateTopicReqPermissionEnum struct {
	SUB CreateTopicReqPermission
	PUB CreateTopicReqPermission
	ALL CreateTopicReqPermission
}

func GetCreateTopicReqPermissionEnum() CreateTopicReqPermissionEnum {
	return CreateTopicReqPermissionEnum{
		SUB: CreateTopicReqPermission{
			value: "sub",
		},
		PUB: CreateTopicReqPermission{
			value: "pub",
		},
		ALL: CreateTopicReqPermission{
			value: "all",
		},
	}
}

func (c CreateTopicReqPermission) Value() string {
	return c.value
}

func (c CreateTopicReqPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTopicReqPermission) UnmarshalJSON(b []byte) error {
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

type CreateTopicReqMessageType struct {
	value string
}

type CreateTopicReqMessageTypeEnum struct {
	NORMAL      CreateTopicReqMessageType
	FIFO        CreateTopicReqMessageType
	DELAY       CreateTopicReqMessageType
	TRANSACTION CreateTopicReqMessageType
}

func GetCreateTopicReqMessageTypeEnum() CreateTopicReqMessageTypeEnum {
	return CreateTopicReqMessageTypeEnum{
		NORMAL: CreateTopicReqMessageType{
			value: "NORMAL",
		},
		FIFO: CreateTopicReqMessageType{
			value: "FIFO",
		},
		DELAY: CreateTopicReqMessageType{
			value: "DELAY",
		},
		TRANSACTION: CreateTopicReqMessageType{
			value: "TRANSACTION",
		},
	}
}

func (c CreateTopicReqMessageType) Value() string {
	return c.value
}

func (c CreateTopicReqMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTopicReqMessageType) UnmarshalJSON(b []byte) error {
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
