package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PrecheckResult 预检查项结果。
type PrecheckResult struct {

	// 检查项。
	Item *string `json:"item,omitempty"`

	// 检查结果。取值： - PASSED：检查通过。 - ALARM：检查告警项。 - FAILED：检查失败。
	Result *PrecheckResultResult `json:"result,omitempty"`

	// 失败原因。
	FailedReason *string `json:"failed_reason,omitempty"`

	// 失败数据。
	Data *string `json:"data,omitempty"`

	// 失败详情。
	RawErrorMsg *string `json:"raw_error_msg,omitempty"`

	// 检查项分组。
	Group *string `json:"group,omitempty"`

	// 是否支持跳过。
	IsSupportSkip *bool `json:"is_support_skip,omitempty"`

	// 是否已跳过。
	IsSkipped *bool `json:"is_skipped,omitempty"`

	// 失败子任务详情。
	FailedSubJobs *[]PrecheckFailSubJobResult `json:"failed_sub_jobs,omitempty"`
}

func (o PrecheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckResult struct{}"
	}

	return strings.Join([]string{"PrecheckResult", string(data)}, " ")
}

type PrecheckResultResult struct {
	value string
}

type PrecheckResultResultEnum struct {
	PASSED PrecheckResultResult
	ALARM  PrecheckResultResult
	FAILED PrecheckResultResult
}

func GetPrecheckResultResultEnum() PrecheckResultResultEnum {
	return PrecheckResultResultEnum{
		PASSED: PrecheckResultResult{
			value: "PASSED",
		},
		ALARM: PrecheckResultResult{
			value: "ALARM",
		},
		FAILED: PrecheckResultResult{
			value: "FAILED",
		},
	}
}

func (c PrecheckResultResult) Value() string {
	return c.value
}

func (c PrecheckResultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrecheckResultResult) UnmarshalJSON(b []byte) error {
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
