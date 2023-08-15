package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyDto 包含有关策略的详细信息的结构。
type PolicyDto struct {

	// 策略的文本内容。
	Content string `json:"content"`

	PolicySummary *PolicySummaryDto `json:"policy_summary"`
}

func (o PolicyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDto struct{}"
	}

	return strings.Join([]string{"PolicyDto", string(data)}, " ")
}
