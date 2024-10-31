package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StatusEnum 运行状态PENDING, CREATING, RUNNING, UPDATING, SUCCEEDED, FAILED, STOPPING, STOPPED, DELETING, DELETED
type StatusEnum struct {
	value string
}

type StatusEnumEnum struct {
	PENDING   StatusEnum
	CREATING  StatusEnum
	RUNNING   StatusEnum
	UPDATING  StatusEnum
	SUCCEEDED StatusEnum
	FAILED    StatusEnum
	STOPPING  StatusEnum
	STOPPED   StatusEnum
	DELETING  StatusEnum
	DELETED   StatusEnum
}

func GetStatusEnumEnum() StatusEnumEnum {
	return StatusEnumEnum{
		PENDING: StatusEnum{
			value: "PENDING",
		},
		CREATING: StatusEnum{
			value: "CREATING",
		},
		RUNNING: StatusEnum{
			value: "RUNNING",
		},
		UPDATING: StatusEnum{
			value: "UPDATING",
		},
		SUCCEEDED: StatusEnum{
			value: "SUCCEEDED",
		},
		FAILED: StatusEnum{
			value: "FAILED",
		},
		STOPPING: StatusEnum{
			value: "STOPPING",
		},
		STOPPED: StatusEnum{
			value: "STOPPED",
		},
		DELETING: StatusEnum{
			value: "DELETING",
		},
		DELETED: StatusEnum{
			value: "DELETED",
		},
	}
}

func (c StatusEnum) Value() string {
	return c.value
}

func (c StatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatusEnum) UnmarshalJSON(b []byte) error {
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
