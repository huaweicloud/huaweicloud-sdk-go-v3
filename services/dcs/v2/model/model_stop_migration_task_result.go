package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StopMigrationTaskResult struct {

	// 下发停止迁移任务操作结果。
	Result *StopMigrationTaskResultResult `json:"result,omitempty"`

	// 数据迁移任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o StopMigrationTaskResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskResult struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskResult", string(data)}, " ")
}

type StopMigrationTaskResultResult struct {
	value string
}

type StopMigrationTaskResultResultEnum struct {
	SUCCESS StopMigrationTaskResultResult
	FAILED  StopMigrationTaskResultResult
}

func GetStopMigrationTaskResultResultEnum() StopMigrationTaskResultResultEnum {
	return StopMigrationTaskResultResultEnum{
		SUCCESS: StopMigrationTaskResultResult{
			value: "success",
		},
		FAILED: StopMigrationTaskResultResult{
			value: "failed",
		},
	}
}

func (c StopMigrationTaskResultResult) Value() string {
	return c.value
}

func (c StopMigrationTaskResultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResultResult) UnmarshalJSON(b []byte) error {
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
