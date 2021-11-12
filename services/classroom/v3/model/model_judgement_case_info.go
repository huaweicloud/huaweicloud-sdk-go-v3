package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用例数据
type JudgementCaseInfo struct {
	// 用例数据输入

	Input string `json:"input"`
	// 用例数据期望输出

	Output *string `json:"output,omitempty"`
}

func (o JudgementCaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JudgementCaseInfo struct{}"
	}

	return strings.Join([]string{"JudgementCaseInfo", string(data)}, " ")
}
