package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteFlinkJobSavepointRequestBody 创建savepoint请求体
type ExecuteFlinkJobSavepointRequestBody struct {

	// 操作类型
	Action ExecuteFlinkJobSavepointRequestBodyAction `json:"action"`

	// obs桶路径.例 \"bucket_name/file_name/\"
	SavepointPath string `json:"savepoint_path"`
}

func (o ExecuteFlinkJobSavepointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteFlinkJobSavepointRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteFlinkJobSavepointRequestBody", string(data)}, " ")
}

type ExecuteFlinkJobSavepointRequestBodyAction struct {
	value string
}

type ExecuteFlinkJobSavepointRequestBodyActionEnum struct {
	TRIGGER ExecuteFlinkJobSavepointRequestBodyAction
}

func GetExecuteFlinkJobSavepointRequestBodyActionEnum() ExecuteFlinkJobSavepointRequestBodyActionEnum {
	return ExecuteFlinkJobSavepointRequestBodyActionEnum{
		TRIGGER: ExecuteFlinkJobSavepointRequestBodyAction{
			value: "TRIGGER",
		},
	}
}

func (c ExecuteFlinkJobSavepointRequestBodyAction) Value() string {
	return c.value
}

func (c ExecuteFlinkJobSavepointRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteFlinkJobSavepointRequestBodyAction) UnmarshalJSON(b []byte) error {
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
