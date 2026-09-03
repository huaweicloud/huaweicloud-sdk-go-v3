package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskErrorPolicy struct {

	// 小网拨测：同一个ip异常的用例大于多少个告警
	SameIpErrorTestCaseCount *int32 `json:"sameIpErrorTestCaseCount,omitempty"`

	// 小网拨测:同一用例在N个IP中异常，并且异常的用例个数达到M个告警
	SameTestCaseErrorIpCount *string `json:"sameTestCaseErrorIpCount,omitempty"`

	// 任务中多少个用例异常告警
	TestCaseErrorCount *int32 `json:"testCaseErrorCount,omitempty"`

	// 任务中多少百分比的用例异常告警
	TestCaseErrorRatio *int32 `json:"testCaseErrorRatio,omitempty"`
}

func (o TaskErrorPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskErrorPolicy struct{}"
	}

	return strings.Join([]string{"TaskErrorPolicy", string(data)}, " ")
}
