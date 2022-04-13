package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type SetOnlineMigrationTaskResponse struct {
	// 迁移任务ID。

	Id *string `json:"id,omitempty"`
	// 迁移任务名称。

	Name *string `json:"name,omitempty"`
	// 迁移任务状态，这个字段的值包括：SUCCESS, FAILED, MIGRATING，TERMINATED

	Status         *SetOnlineMigrationTaskResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o SetOnlineMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOnlineMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"SetOnlineMigrationTaskResponse", string(data)}, " ")
}

type SetOnlineMigrationTaskResponseStatus struct {
	value string
}

type SetOnlineMigrationTaskResponseStatusEnum struct {
	SUCCESS    SetOnlineMigrationTaskResponseStatus
	FAILED     SetOnlineMigrationTaskResponseStatus
	MIGRATING  SetOnlineMigrationTaskResponseStatus
	TERMINATED SetOnlineMigrationTaskResponseStatus
}

func GetSetOnlineMigrationTaskResponseStatusEnum() SetOnlineMigrationTaskResponseStatusEnum {
	return SetOnlineMigrationTaskResponseStatusEnum{
		SUCCESS: SetOnlineMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: SetOnlineMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: SetOnlineMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: SetOnlineMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c SetOnlineMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetOnlineMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
