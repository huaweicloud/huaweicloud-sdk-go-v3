package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AutoOperationConfig 资产自动处理任务配置。
type AutoOperationConfig struct {

	// BLOCK: 冻结 DELETE：删除
	Operation *AutoOperationConfigOperation `json:"operation,omitempty"`

	// 资源过期时间，格式遵循：RFC 3339 如\"2025-01-10T00:00:00Z\"
	OperationTime *string `json:"operation_time,omitempty"`
}

func (o AutoOperationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoOperationConfig struct{}"
	}

	return strings.Join([]string{"AutoOperationConfig", string(data)}, " ")
}

type AutoOperationConfigOperation struct {
	value string
}

type AutoOperationConfigOperationEnum struct {
	BLOCK  AutoOperationConfigOperation
	DELETE AutoOperationConfigOperation
}

func GetAutoOperationConfigOperationEnum() AutoOperationConfigOperationEnum {
	return AutoOperationConfigOperationEnum{
		BLOCK: AutoOperationConfigOperation{
			value: "BLOCK",
		},
		DELETE: AutoOperationConfigOperation{
			value: "DELETE",
		},
	}
}

func (c AutoOperationConfigOperation) Value() string {
	return c.value
}

func (c AutoOperationConfigOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoOperationConfigOperation) UnmarshalJSON(b []byte) error {
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
