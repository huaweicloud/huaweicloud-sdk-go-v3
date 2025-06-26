package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteRetentionPolicyRequestBody struct {

	// 是否模拟运行
	DryRun bool `json:"dry_run"`
}

func (o ExecuteRetentionPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteRetentionPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteRetentionPolicyRequestBody", string(data)}, " ")
}
