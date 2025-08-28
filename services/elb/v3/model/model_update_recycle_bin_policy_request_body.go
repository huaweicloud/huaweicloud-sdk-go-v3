package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinPolicyRequestBody **参数解释**：更新回收站配置的请求体。
type UpdateRecycleBinPolicyRequestBody struct {
	RecycleBin *RecycleBinPolicyRequestBody `json:"recycle_bin,omitempty"`
}

func (o UpdateRecycleBinPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinPolicyRequestBody", string(data)}, " ")
}
