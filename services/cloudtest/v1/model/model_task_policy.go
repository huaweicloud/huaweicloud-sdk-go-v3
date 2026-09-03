package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskPolicy struct {

	// 任务连续失败N次告警
	FailedTimes *int32 `json:"failed_times,omitempty"`

	// 小网拨测：同一个ip失败的用例大于多少个告警
	SameIpFailedTestCaseCount *int32 `json:"sameIpFailedTestCaseCount,omitempty"`

	// 小网拨测:同一用例在N个IP中失败，并且失败的用例个数达到M个
	SameTestCaseFailedIpCount *string `json:"sameTestCaseFailedIpCount,omitempty"`

	// 任务中多少个用例失败告警
	TestCaseFailedCount *int32 `json:"testCaseFailedCount,omitempty"`

	// 任务中多少百分比的用例失败告警
	TestCaseFailedRatio *int32 `json:"testCaseFailedRatio,omitempty"`
}

func (o TaskPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskPolicy struct{}"
	}

	return strings.Join([]string{"TaskPolicy", string(data)}, " ")
}
