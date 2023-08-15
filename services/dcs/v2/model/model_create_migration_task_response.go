package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateMigrationTaskResponse Response Object
type CreateMigrationTaskResponse struct {

	// 迁移任务ID。
	Id *string `json:"id,omitempty"`

	// 迁移任务名称。
	Name *string `json:"name,omitempty"`

	// 迁移任务状态，这个字段的值包括：SUCCESS, FAILED, MIGRATING，TERMINATED
	Status         *CreateMigrationTaskResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o CreateMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateMigrationTaskResponse", string(data)}, " ")
}

type CreateMigrationTaskResponseStatus struct {
	value string
}

type CreateMigrationTaskResponseStatusEnum struct {
	SUCCESS    CreateMigrationTaskResponseStatus
	FAILED     CreateMigrationTaskResponseStatus
	MIGRATING  CreateMigrationTaskResponseStatus
	TERMINATED CreateMigrationTaskResponseStatus
}

func GetCreateMigrationTaskResponseStatusEnum() CreateMigrationTaskResponseStatusEnum {
	return CreateMigrationTaskResponseStatusEnum{
		SUCCESS: CreateMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: CreateMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: CreateMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: CreateMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c CreateMigrationTaskResponseStatus) Value() string {
	return c.value
}

func (c CreateMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
