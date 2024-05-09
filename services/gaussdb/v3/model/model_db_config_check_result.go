package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DbConfigCheckResult 库配置校验检查结果。
type DbConfigCheckResult struct {

	// 参数名。
	ParamName *string `json:"param_name,omitempty"`

	// 参数值。
	Value *string `json:"value,omitempty"`

	// 校验结果。
	CheckResult *DbConfigCheckResultCheckResult `json:"check_result,omitempty"`
}

func (o DbConfigCheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbConfigCheckResult struct{}"
	}

	return strings.Join([]string{"DbConfigCheckResult", string(data)}, " ")
}

type DbConfigCheckResultCheckResult struct {
	value string
}

type DbConfigCheckResultCheckResultEnum struct {
	SUCCESS DbConfigCheckResultCheckResult
	FAIL    DbConfigCheckResultCheckResult
}

func GetDbConfigCheckResultCheckResultEnum() DbConfigCheckResultCheckResultEnum {
	return DbConfigCheckResultCheckResultEnum{
		SUCCESS: DbConfigCheckResultCheckResult{
			value: "success",
		},
		FAIL: DbConfigCheckResultCheckResult{
			value: "fail",
		},
	}
}

func (c DbConfigCheckResultCheckResult) Value() string {
	return c.value
}

func (c DbConfigCheckResultCheckResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DbConfigCheckResultCheckResult) UnmarshalJSON(b []byte) error {
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
