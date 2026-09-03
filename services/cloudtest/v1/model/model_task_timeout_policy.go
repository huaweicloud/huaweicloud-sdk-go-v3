package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskTimeoutPolicy struct {

	// 小网拨测：同一个ip超时的用例大于多少个告警
	SameIpTimeoutTestCaseCount *int32 `json:"sameIpTimeoutTestCaseCount,omitempty"`

	// 小网拨测:同一用例在N个IP中超时，并且超时的用例个数达到M个告警
	SameTestCaseTimeoutIpCount *string `json:"sameTestCaseTimeoutIpCount,omitempty"`

	// 任务中多少个用例超时告警
	TestCaseTimeoutCount *int32 `json:"testCaseTimeoutCount,omitempty"`

	// 任务中多少百分比的用例超时告警
	TestCaseTimeoutRatio *int32 `json:"testCaseTimeoutRatio,omitempty"`

	// 任务连续超时告警
	TimeoutTimes *int32 `json:"timeoutTimes,omitempty"`
}

func (o TaskTimeoutPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskTimeoutPolicy struct{}"
	}

	return strings.Join([]string{"TaskTimeoutPolicy", string(data)}, " ")
}
