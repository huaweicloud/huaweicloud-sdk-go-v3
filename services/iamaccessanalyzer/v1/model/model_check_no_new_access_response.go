package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckNoNewAccessResponse Response Object
type CheckNoNewAccessResponse struct {

	// 更新后的策略是否允许新访问权限的消息。
	Message *string `json:"message,omitempty"`

	// 检查新访问权限的结果。
	CheckResult *CheckNoNewAccessResponseCheckResult `json:"check_result,omitempty"`

	// 新增action的statement描述。
	Reasons        *[]CheckNoNewAccessReason `json:"reasons,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o CheckNoNewAccessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNoNewAccessResponse struct{}"
	}

	return strings.Join([]string{"CheckNoNewAccessResponse", string(data)}, " ")
}

type CheckNoNewAccessResponseCheckResult struct {
	value string
}

type CheckNoNewAccessResponseCheckResultEnum struct {
	PASS CheckNoNewAccessResponseCheckResult
	FAIL CheckNoNewAccessResponseCheckResult
}

func GetCheckNoNewAccessResponseCheckResultEnum() CheckNoNewAccessResponseCheckResultEnum {
	return CheckNoNewAccessResponseCheckResultEnum{
		PASS: CheckNoNewAccessResponseCheckResult{
			value: "pass",
		},
		FAIL: CheckNoNewAccessResponseCheckResult{
			value: "fail",
		},
	}
}

func (c CheckNoNewAccessResponseCheckResult) Value() string {
	return c.value
}

func (c CheckNoNewAccessResponseCheckResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckNoNewAccessResponseCheckResult) UnmarshalJSON(b []byte) error {
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
