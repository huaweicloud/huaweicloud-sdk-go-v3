package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAuraDataframeJobRequest Request Object
type CancelAuraDataframeJobRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：作业ID。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	JobId string `json:"job_id"`
}

func (o CancelAuraDataframeJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAuraDataframeJobRequest struct{}"
	}

	return strings.Join([]string{"CancelAuraDataframeJobRequest", string(data)}, " ")
}
