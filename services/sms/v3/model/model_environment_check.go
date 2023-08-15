package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EnvironmentCheck 源端校验项
type EnvironmentCheck struct {

	// 该检查项的ID
	Id int64 `json:"id"`

	// 参数
	Params *[]string `json:"params,omitempty"`

	// 检查项名称
	Name string `json:"name"`

	// 检查结果 OK：检查通过 WARN：警告 ERROR:检查不通过
	Result EnvironmentCheckResult `json:"result"`

	// 检查不通过的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 检查的错误或者警告
	ErrorOrWarn *string `json:"error_or_warn,omitempty"`

	// 检查不通过的错误参数
	ErrorParams *string `json:"error_params,omitempty"`
}

func (o EnvironmentCheck) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentCheck struct{}"
	}

	return strings.Join([]string{"EnvironmentCheck", string(data)}, " ")
}

type EnvironmentCheckResult struct {
	value string
}

type EnvironmentCheckResultEnum struct {
	OK    EnvironmentCheckResult
	WARN  EnvironmentCheckResult
	ERROR EnvironmentCheckResult
}

func GetEnvironmentCheckResultEnum() EnvironmentCheckResultEnum {
	return EnvironmentCheckResultEnum{
		OK: EnvironmentCheckResult{
			value: "OK",
		},
		WARN: EnvironmentCheckResult{
			value: "WARN",
		},
		ERROR: EnvironmentCheckResult{
			value: "ERROR",
		},
	}
}

func (c EnvironmentCheckResult) Value() string {
	return c.value
}

func (c EnvironmentCheckResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvironmentCheckResult) UnmarshalJSON(b []byte) error {
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
