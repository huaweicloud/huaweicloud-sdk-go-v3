package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PreTestCaseInfo struct {
	AlertTemplate *AlertTemplate `json:"alert_template,omitempty"`

	// 0 关闭，1开启
	Enable *string `json:"enable,omitempty"`

	// 用例列表
	TestCases *[]TestCaseBasicInfo `json:"testCases,omitempty"`
}

func (o PreTestCaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreTestCaseInfo struct{}"
	}

	return strings.Join([]string{"PreTestCaseInfo", string(data)}, " ")
}
