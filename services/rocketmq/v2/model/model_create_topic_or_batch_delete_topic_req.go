package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTopicOrBatchDeleteTopicReq struct {

	// 主题名称，只能由英文字母、数字、百分号、竖线、中划线、下划线组成，长度3~64个字符。
	Name *string `json:"name,omitempty"`

	// 关联的代理。
	Brokers *[]string `json:"brokers,omitempty"`

	// 队列数，范围1~50。
	QueueNum float32 `json:"queue_num,omitempty"`

	// 权限。
	Permission *CreateTopicOrBatchDeleteTopicReqPermission `json:"permission,omitempty"`

	// 消息类型（RocketMQ实例5.x版本才包含此参数）。
	MessageType *CreateTopicOrBatchDeleteTopicReqMessageType `json:"message_type,omitempty"`

	// 主题列表，当批量删除主题时使用。
	Topics *[]string `json:"topics,omitempty"`
}

func (o CreateTopicOrBatchDeleteTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicOrBatchDeleteTopicReq struct{}"
	}

	return strings.Join([]string{"CreateTopicOrBatchDeleteTopicReq", string(data)}, " ")
}

type CreateTopicOrBatchDeleteTopicReqPermission struct {
	value string
}

type CreateTopicOrBatchDeleteTopicReqPermissionEnum struct {
	SUB CreateTopicOrBatchDeleteTopicReqPermission
	PUB CreateTopicOrBatchDeleteTopicReqPermission
	ALL CreateTopicOrBatchDeleteTopicReqPermission
}

func GetCreateTopicOrBatchDeleteTopicReqPermissionEnum() CreateTopicOrBatchDeleteTopicReqPermissionEnum {
	return CreateTopicOrBatchDeleteTopicReqPermissionEnum{
		SUB: CreateTopicOrBatchDeleteTopicReqPermission{
			value: "sub",
		},
		PUB: CreateTopicOrBatchDeleteTopicReqPermission{
			value: "pub",
		},
		ALL: CreateTopicOrBatchDeleteTopicReqPermission{
			value: "all",
		},
	}
}

func (c CreateTopicOrBatchDeleteTopicReqPermission) Value() string {
	return c.value
}

func (c CreateTopicOrBatchDeleteTopicReqPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTopicOrBatchDeleteTopicReqPermission) UnmarshalJSON(b []byte) error {
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

type CreateTopicOrBatchDeleteTopicReqMessageType struct {
	value string
}

type CreateTopicOrBatchDeleteTopicReqMessageTypeEnum struct {
	NORMAL      CreateTopicOrBatchDeleteTopicReqMessageType
	FIFO        CreateTopicOrBatchDeleteTopicReqMessageType
	DELAY       CreateTopicOrBatchDeleteTopicReqMessageType
	TRANSACTION CreateTopicOrBatchDeleteTopicReqMessageType
}

func GetCreateTopicOrBatchDeleteTopicReqMessageTypeEnum() CreateTopicOrBatchDeleteTopicReqMessageTypeEnum {
	return CreateTopicOrBatchDeleteTopicReqMessageTypeEnum{
		NORMAL: CreateTopicOrBatchDeleteTopicReqMessageType{
			value: "NORMAL",
		},
		FIFO: CreateTopicOrBatchDeleteTopicReqMessageType{
			value: "FIFO",
		},
		DELAY: CreateTopicOrBatchDeleteTopicReqMessageType{
			value: "DELAY",
		},
		TRANSACTION: CreateTopicOrBatchDeleteTopicReqMessageType{
			value: "TRANSACTION",
		},
	}
}

func (c CreateTopicOrBatchDeleteTopicReqMessageType) Value() string {
	return c.value
}

func (c CreateTopicOrBatchDeleteTopicReqMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTopicOrBatchDeleteTopicReqMessageType) UnmarshalJSON(b []byte) error {
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
