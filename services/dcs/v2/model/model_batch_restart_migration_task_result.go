package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchRestartMigrationTaskResult struct {

	// 下发重启迁移任务操作结果。
	Result *BatchRestartMigrationTaskResultResult `json:"result,omitempty"`

	// 数据迁移任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o BatchRestartMigrationTaskResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartMigrationTaskResult struct{}"
	}

	return strings.Join([]string{"BatchRestartMigrationTaskResult", string(data)}, " ")
}

type BatchRestartMigrationTaskResultResult struct {
	value string
}

type BatchRestartMigrationTaskResultResultEnum struct {
	SUCCESS BatchRestartMigrationTaskResultResult
	FAILED  BatchRestartMigrationTaskResultResult
}

func GetBatchRestartMigrationTaskResultResultEnum() BatchRestartMigrationTaskResultResultEnum {
	return BatchRestartMigrationTaskResultResultEnum{
		SUCCESS: BatchRestartMigrationTaskResultResult{
			value: "success",
		},
		FAILED: BatchRestartMigrationTaskResultResult{
			value: "failed",
		},
	}
}

func (c BatchRestartMigrationTaskResultResult) Value() string {
	return c.value
}

func (c BatchRestartMigrationTaskResultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRestartMigrationTaskResultResult) UnmarshalJSON(b []byte) error {
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
