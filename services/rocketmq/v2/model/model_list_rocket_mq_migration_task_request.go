package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRocketMqMigrationTaskRequest Request Object
type ListRocketMqMigrationTaskRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 查询类型。 **约束限制**： 不涉及。 **取值范围**： - vhost：虚拟主机。 - exchange：交换机。 - queue：队列。 - all：所有。            **默认取值**： 不涉及。
	Type *ListRocketMqMigrationTaskRequestType `json:"type,omitempty"`

	// **参数解释**： 当前页，从1开始。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 当前页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： vhost名称 **约束限制**： - 查询vhost列表时，该字段可为空。 - 查询exchange列表时，该字段为exchange所属vhost名称。 - 查询queue列表时，该字段为queue所属vhost-所属exchange，例vhost1-exchange1。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ListRocketMqMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRocketMqMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"ListRocketMqMigrationTaskRequest", string(data)}, " ")
}

type ListRocketMqMigrationTaskRequestType struct {
	value string
}

type ListRocketMqMigrationTaskRequestTypeEnum struct {
	VHOST    ListRocketMqMigrationTaskRequestType
	EXCHANGE ListRocketMqMigrationTaskRequestType
	QUEUE    ListRocketMqMigrationTaskRequestType
	ALL      ListRocketMqMigrationTaskRequestType
}

func GetListRocketMqMigrationTaskRequestTypeEnum() ListRocketMqMigrationTaskRequestTypeEnum {
	return ListRocketMqMigrationTaskRequestTypeEnum{
		VHOST: ListRocketMqMigrationTaskRequestType{
			value: "vhost",
		},
		EXCHANGE: ListRocketMqMigrationTaskRequestType{
			value: "exchange",
		},
		QUEUE: ListRocketMqMigrationTaskRequestType{
			value: "queue",
		},
		ALL: ListRocketMqMigrationTaskRequestType{
			value: "all",
		},
	}
}

func (c ListRocketMqMigrationTaskRequestType) Value() string {
	return c.value
}

func (c ListRocketMqMigrationTaskRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRocketMqMigrationTaskRequestType) UnmarshalJSON(b []byte) error {
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
