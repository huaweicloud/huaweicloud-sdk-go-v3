package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NotebookStatus struct {
	value string
}

type NotebookStatusEnum struct {
	RUNNING        NotebookStatus
	ABNORMAL       NotebookStatus
	HIBERNATE      NotebookStatus
	SUCCEEDED      NotebookStatus
	CREATING       NotebookStatus
	DELETING       NotebookStatus
	UPDATING       NotebookStatus
	CREATED_FAILED NotebookStatus
	DELETED_FAILED NotebookStatus
	UPDATED_FAILED NotebookStatus
	UNKNOWN        NotebookStatus
}

func GetNotebookStatusEnum() NotebookStatusEnum {
	return NotebookStatusEnum{
		RUNNING: NotebookStatus{
			value: "Running",
		},
		ABNORMAL: NotebookStatus{
			value: "Abnormal",
		},
		HIBERNATE: NotebookStatus{
			value: "Hibernate",
		},
		SUCCEEDED: NotebookStatus{
			value: "Succeeded",
		},
		CREATING: NotebookStatus{
			value: "Creating",
		},
		DELETING: NotebookStatus{
			value: "Deleting",
		},
		UPDATING: NotebookStatus{
			value: "Updating",
		},
		CREATED_FAILED: NotebookStatus{
			value: "CreatedFailed",
		},
		DELETED_FAILED: NotebookStatus{
			value: "DeletedFailed",
		},
		UPDATED_FAILED: NotebookStatus{
			value: "UpdatedFailed",
		},
		UNKNOWN: NotebookStatus{
			value: "Unknown",
		},
	}
}

func (c NotebookStatus) Value() string {
	return c.value
}

func (c NotebookStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotebookStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
