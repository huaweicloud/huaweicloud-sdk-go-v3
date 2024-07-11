package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRocketMqMigrationTaskRequest Request Object
type ListRocketMqMigrationTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 查询类型。
	Type *ListRocketMqMigrationTaskRequestType `json:"type,omitempty"`

	// 当前页，从1开始。
	Offset *string `json:"offset,omitempty"`

	// 当前页大小。
	Limit *string `json:"limit,omitempty"`

	// - 查询vhost列表时，该字段可为空。 - 查询exchange列表时，该字段为exchange所属vhost名称。 - 查询queue列表时，该字段为queue所属vhost-所属exchange，例vhost1-exchange1。
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
