package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRocketMqMigrationTaskRequest Request Object
type CreateRocketMqMigrationTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// true开启同名覆盖，会对已有的同名元数据的配置进行修改，false时当Topic或group已存在则会报错。
	Overwrite CreateRocketMqMigrationTaskRequestOverwrite `json:"overwrite"`

	// 迁移任务名称，名称规则参考创建实例。
	Name string `json:"name"`

	// 迁移任务类型，分为自建RocketMQ上云(rocketmq)、自建RabbitMQ上云(rabbitToRocket)
	Type CreateRocketMqMigrationTaskRequestType `json:"type"`

	Body *CreateRocketMqMigrationTaskReq `json:"body,omitempty"`
}

func (o CreateRocketMqMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRocketMqMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateRocketMqMigrationTaskRequest", string(data)}, " ")
}

type CreateRocketMqMigrationTaskRequestOverwrite struct {
	value string
}

type CreateRocketMqMigrationTaskRequestOverwriteEnum struct {
	TRUE  CreateRocketMqMigrationTaskRequestOverwrite
	FALSE CreateRocketMqMigrationTaskRequestOverwrite
}

func GetCreateRocketMqMigrationTaskRequestOverwriteEnum() CreateRocketMqMigrationTaskRequestOverwriteEnum {
	return CreateRocketMqMigrationTaskRequestOverwriteEnum{
		TRUE: CreateRocketMqMigrationTaskRequestOverwrite{
			value: "true",
		},
		FALSE: CreateRocketMqMigrationTaskRequestOverwrite{
			value: "false",
		},
	}
}

func (c CreateRocketMqMigrationTaskRequestOverwrite) Value() string {
	return c.value
}

func (c CreateRocketMqMigrationTaskRequestOverwrite) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRocketMqMigrationTaskRequestOverwrite) UnmarshalJSON(b []byte) error {
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

type CreateRocketMqMigrationTaskRequestType struct {
	value string
}

type CreateRocketMqMigrationTaskRequestTypeEnum struct {
	ROCKETMQ         CreateRocketMqMigrationTaskRequestType
	RABBIT_TO_ROCKET CreateRocketMqMigrationTaskRequestType
}

func GetCreateRocketMqMigrationTaskRequestTypeEnum() CreateRocketMqMigrationTaskRequestTypeEnum {
	return CreateRocketMqMigrationTaskRequestTypeEnum{
		ROCKETMQ: CreateRocketMqMigrationTaskRequestType{
			value: "rocketmq",
		},
		RABBIT_TO_ROCKET: CreateRocketMqMigrationTaskRequestType{
			value: "rabbitToRocket",
		},
	}
}

func (c CreateRocketMqMigrationTaskRequestType) Value() string {
	return c.value
}

func (c CreateRocketMqMigrationTaskRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRocketMqMigrationTaskRequestType) UnmarshalJSON(b []byte) error {
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
