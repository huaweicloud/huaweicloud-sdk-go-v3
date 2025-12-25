package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetryPolicyRequestBodyDataObjectBlockAge 阻断老化
type CreateRetryPolicyRequestBodyDataObjectBlockAge struct {

	// 是否阻断老化
	IsBlockAgeing bool `json:"is_block_ageing"`

	// 老化时间，毫秒级时间戳
	BlockAgeing *string `json:"block_ageing,omitempty"`
}

func (o CreateRetryPolicyRequestBodyDataObjectBlockAge) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryPolicyRequestBodyDataObjectBlockAge struct{}"
	}

	return strings.Join([]string{"CreateRetryPolicyRequestBodyDataObjectBlockAge", string(data)}, " ")
}
