package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DiagnoseResult 检测结果 * NO_RISK 无风险 * MEDIUM_RISK 中风险 * HIGH_RISK 高风险 * NOT_SCANNED 未扫描
type DiagnoseResult struct {
	value string
}

type DiagnoseResultEnum struct {
	NO_RISK     DiagnoseResult
	MEDIUM_RISK DiagnoseResult
	HIGH_RISK   DiagnoseResult
	NOT_SCANNED DiagnoseResult
}

func GetDiagnoseResultEnum() DiagnoseResultEnum {
	return DiagnoseResultEnum{
		NO_RISK: DiagnoseResult{
			value: "NO_RISK",
		},
		MEDIUM_RISK: DiagnoseResult{
			value: "MEDIUM_RISK",
		},
		HIGH_RISK: DiagnoseResult{
			value: "HIGH_RISK",
		},
		NOT_SCANNED: DiagnoseResult{
			value: "NOT_SCANNED",
		},
	}
}

func (c DiagnoseResult) Value() string {
	return c.value
}

func (c DiagnoseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnoseResult) UnmarshalJSON(b []byte) error {
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
