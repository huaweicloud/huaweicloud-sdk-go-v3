package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOneTopicResponse Response Object
type ShowOneTopicResponse struct {

	// **参数解释**： Topic名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 总读队列个数。 **取值范围**： 不涉及。
	TotalReadQueueNum float32 `json:"total_read_queue_num,omitempty"`

	// **参数解释**： 总写队列个数。 **取值范围**： 不涉及。
	TotalWriteQueueNum float32 `json:"total_write_queue_num,omitempty"`

	// **参数解释**： 权限。 **取值范围**： - sub：拥有订阅权限。 - pub：拥有发布权限。 - all：拥有发布、订阅权限。
	Permission *ShowOneTopicResponsePermission `json:"permission,omitempty"`

	// 关联的代理。
	Brokers *[]TopicBrokers `json:"brokers,omitempty"`

	// **参数解释**： 消息类型（RocketMQ实例5.x版本才包含此参数）。 **取值范围**： - NORMAL：普通消息。 - FIFO：顺序消息。 - DELAY：定时消息。 - TRANSACTION：事务消息。
	MessageType    *ShowOneTopicResponseMessageType `json:"message_type,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowOneTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOneTopicResponse struct{}"
	}

	return strings.Join([]string{"ShowOneTopicResponse", string(data)}, " ")
}

type ShowOneTopicResponsePermission struct {
	value string
}

type ShowOneTopicResponsePermissionEnum struct {
	SUB ShowOneTopicResponsePermission
	PUB ShowOneTopicResponsePermission
	ALL ShowOneTopicResponsePermission
}

func GetShowOneTopicResponsePermissionEnum() ShowOneTopicResponsePermissionEnum {
	return ShowOneTopicResponsePermissionEnum{
		SUB: ShowOneTopicResponsePermission{
			value: "sub",
		},
		PUB: ShowOneTopicResponsePermission{
			value: "pub",
		},
		ALL: ShowOneTopicResponsePermission{
			value: "all",
		},
	}
}

func (c ShowOneTopicResponsePermission) Value() string {
	return c.value
}

func (c ShowOneTopicResponsePermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOneTopicResponsePermission) UnmarshalJSON(b []byte) error {
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

type ShowOneTopicResponseMessageType struct {
	value string
}

type ShowOneTopicResponseMessageTypeEnum struct {
	NORMAL      ShowOneTopicResponseMessageType
	FIFO        ShowOneTopicResponseMessageType
	DELAY       ShowOneTopicResponseMessageType
	TRANSACTION ShowOneTopicResponseMessageType
}

func GetShowOneTopicResponseMessageTypeEnum() ShowOneTopicResponseMessageTypeEnum {
	return ShowOneTopicResponseMessageTypeEnum{
		NORMAL: ShowOneTopicResponseMessageType{
			value: "NORMAL",
		},
		FIFO: ShowOneTopicResponseMessageType{
			value: "FIFO",
		},
		DELAY: ShowOneTopicResponseMessageType{
			value: "DELAY",
		},
		TRANSACTION: ShowOneTopicResponseMessageType{
			value: "TRANSACTION",
		},
	}
}

func (c ShowOneTopicResponseMessageType) Value() string {
	return c.value
}

func (c ShowOneTopicResponseMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOneTopicResponseMessageType) UnmarshalJSON(b []byte) error {
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
