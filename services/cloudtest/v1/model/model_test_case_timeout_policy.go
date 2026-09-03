package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestCaseTimeoutPolicy struct {

	// 用例超时多少次告警
	TimeoutTimes *int32 `json:"timeoutTimes,omitempty"`
}

func (o TestCaseTimeoutPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseTimeoutPolicy struct{}"
	}

	return strings.Join([]string{"TestCaseTimeoutPolicy", string(data)}, " ")
}
