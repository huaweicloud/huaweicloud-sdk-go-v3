package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskTypeEnum 任务类型枚举
type TaskTypeEnum struct {
	value string
}

type TaskTypeEnumEnum struct {
	BACKUP          TaskTypeEnum
	REPLICATION     TaskTypeEnum
	RESTORE         TaskTypeEnum
	DELETE          TaskTypeEnum
	VAULT_DELETE    TaskTypeEnum
	REMOVE_RESOURCE TaskTypeEnum
}

func GetTaskTypeEnumEnum() TaskTypeEnumEnum {
	return TaskTypeEnumEnum{
		BACKUP: TaskTypeEnum{
			value: "backup",
		},
		REPLICATION: TaskTypeEnum{
			value: "replication",
		},
		RESTORE: TaskTypeEnum{
			value: "restore",
		},
		DELETE: TaskTypeEnum{
			value: "delete",
		},
		VAULT_DELETE: TaskTypeEnum{
			value: "vault_delete",
		},
		REMOVE_RESOURCE: TaskTypeEnum{
			value: "remove_resource",
		},
	}
}

func (c TaskTypeEnum) Value() string {
	return c.value
}

func (c TaskTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskTypeEnum) UnmarshalJSON(b []byte) error {
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
