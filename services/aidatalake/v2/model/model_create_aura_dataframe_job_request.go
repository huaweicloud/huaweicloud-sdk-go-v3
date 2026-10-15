package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuraDataframeJobRequest Request Object
type CreateAuraDataframeJobRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	Body *CreateDataframeJobRequestBody `json:"body,omitempty"`
}

func (o CreateAuraDataframeJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuraDataframeJobRequest struct{}"
	}

	return strings.Join([]string{"CreateAuraDataframeJobRequest", string(data)}, " ")
}
