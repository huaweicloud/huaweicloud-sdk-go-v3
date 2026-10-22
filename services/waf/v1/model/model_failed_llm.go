package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedLlm struct {

	// 自定义大模型ID
	LlmId *string `json:"llm_id,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`
}

func (o FailedLlm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedLlm struct{}"
	}

	return strings.Join([]string{"FailedLlm", string(data)}, " ")
}
