package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstanceStatusType 实例状态。
type InstanceStatusType struct {
	value string
}

type InstanceStatusTypeEnum struct {
	INITIALIZING     InstanceStatusType
	UPGRADING        InstanceStatusType
	FAILED           InstanceStatusType
	RUNNING          InstanceStatusType
	DOWN             InstanceStatusType
	DELETING         InstanceStatusType
	DELETED          InstanceStatusType
	RESERVED         InstanceStatusType
	STARTING         InstanceStatusType
	STOPPING         InstanceStatusType
	STOPPED          InstanceStatusType
	RESTARTING       InstanceStatusType
	PENDING          InstanceStatusType
	UNKNOWN          InstanceStatusType
	PARTIALLY_FAILED InstanceStatusType
}

func GetInstanceStatusTypeEnum() InstanceStatusTypeEnum {
	return InstanceStatusTypeEnum{
		INITIALIZING: InstanceStatusType{
			value: "INITIALIZING",
		},
		UPGRADING: InstanceStatusType{
			value: "UPGRADING",
		},
		FAILED: InstanceStatusType{
			value: "FAILED",
		},
		RUNNING: InstanceStatusType{
			value: "RUNNING",
		},
		DOWN: InstanceStatusType{
			value: "DOWN",
		},
		DELETING: InstanceStatusType{
			value: "DELETING",
		},
		DELETED: InstanceStatusType{
			value: "DELETED",
		},
		RESERVED: InstanceStatusType{
			value: "RESERVED",
		},
		STARTING: InstanceStatusType{
			value: "STARTING",
		},
		STOPPING: InstanceStatusType{
			value: "STOPPING",
		},
		STOPPED: InstanceStatusType{
			value: "STOPPED",
		},
		RESTARTING: InstanceStatusType{
			value: "RESTARTING",
		},
		PENDING: InstanceStatusType{
			value: "PENDING",
		},
		UNKNOWN: InstanceStatusType{
			value: "UNKNOWN",
		},
		PARTIALLY_FAILED: InstanceStatusType{
			value: "PARTIALLY_FAILED",
		},
	}
}

func (c InstanceStatusType) Value() string {
	return c.value
}

func (c InstanceStatusType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceStatusType) UnmarshalJSON(b []byte) error {
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
