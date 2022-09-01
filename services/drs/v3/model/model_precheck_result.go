package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 预检查结果信息体
type PrecheckResult struct {

	// 检查项。
	Item *string `json:"item,omitempty" xml:"item"`

	// 检查结果
	Result *PrecheckResultResult `json:"result,omitempty" xml:"result"`

	// 失败原因。
	FailedReason *string `json:"failed_reason,omitempty" xml:"failed_reason"`

	// 加密的数据。
	Data *string `json:"data,omitempty" xml:"data"`

	// 行错误信息。
	RawErrorMsg *string `json:"raw_error_msg,omitempty" xml:"raw_error_msg"`

	// 检查项分组
	Group *string `json:"group,omitempty" xml:"group"`

	// 失败的子任务信息。
	FailedSubJobs *[]PrecheckFailSubJobVo `json:"failed_sub_jobs,omitempty" xml:"failed_sub_jobs"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
