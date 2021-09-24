package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 用例运行结果信息
type JudgementCaseResult struct {
	// 用例实际运行结果输出

	Output string `json:"output"`
	// 用例运行结果状态： judgeout判题类型对应：pass（用例比对成功）、failed（用例比对失败）； caseout判题类型对应：success（用例运行成功）、error（用例运行失败）；run_timeout（用例运行超时）

	CaseStatus JudgementCaseResultCaseStatus `json:"case_status"`
}

func (o JudgementCaseResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JudgementCaseResult struct{}"
	}

	return strings.Join([]string{"JudgementCaseResult", string(data)}, " ")
}

type JudgementCaseResultCaseStatus struct {
	value string
}

type JudgementCaseResultCaseStatusEnum struct {
	PASS        JudgementCaseResultCaseStatus
	FAILED      JudgementCaseResultCaseStatus
	SUCCESS     JudgementCaseResultCaseStatus
	ERROR       JudgementCaseResultCaseStatus
	RUN_TIMEOUT JudgementCaseResultCaseStatus
}

func GetJudgementCaseResultCaseStatusEnum() JudgementCaseResultCaseStatusEnum {
	return JudgementCaseResultCaseStatusEnum{
		PASS: JudgementCaseResultCaseStatus{
			value: "pass",
		},
		FAILED: JudgementCaseResultCaseStatus{
			value: "failed",
		},
		SUCCESS: JudgementCaseResultCaseStatus{
			value: "success",
		},
		ERROR: JudgementCaseResultCaseStatus{
			value: "error",
		},
		RUN_TIMEOUT: JudgementCaseResultCaseStatus{
			value: "run_timeout",
		},
	}
}

func (c JudgementCaseResultCaseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *JudgementCaseResultCaseStatus) UnmarshalJSON(b []byte) error {
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
