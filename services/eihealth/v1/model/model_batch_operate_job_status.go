package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchOperateJobStatus 批量操作作业的结果
type BatchOperateJobStatus struct {
	value string
}

type BatchOperateJobStatusEnum struct {
	SUCCEEDED BatchOperateJobStatus
	FAILED    BatchOperateJobStatus
}

func GetBatchOperateJobStatusEnum() BatchOperateJobStatusEnum {
	return BatchOperateJobStatusEnum{
		SUCCEEDED: BatchOperateJobStatus{
			value: "SUCCEEDED",
		},
		FAILED: BatchOperateJobStatus{
			value: "FAILED",
		},
	}
}

func (c BatchOperateJobStatus) Value() string {
	return c.value
}

func (c BatchOperateJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchOperateJobStatus) UnmarshalJSON(b []byte) error {
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
