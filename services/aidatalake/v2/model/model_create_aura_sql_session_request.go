package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuraSqlSessionRequest Request Object
type CreateAuraSqlSessionRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	Body *CreateSessionRequestBodyV2 `json:"body,omitempty"`
}

func (o CreateAuraSqlSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuraSqlSessionRequest struct{}"
	}

	return strings.Join([]string{"CreateAuraSqlSessionRequest", string(data)}, " ")
}
