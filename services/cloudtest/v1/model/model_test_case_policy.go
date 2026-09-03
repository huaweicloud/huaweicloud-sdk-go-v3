package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestCasePolicy struct {

	// 单用例失败多少次告警
	FailedTimes *int32 `json:"failed_times,omitempty"`

	// 单用例重试多少次后告警
	RetryTimes *int32 `json:"retryTimes,omitempty"`
}

func (o TestCasePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCasePolicy struct{}"
	}

	return strings.Join([]string{"TestCasePolicy", string(data)}, " ")
}
