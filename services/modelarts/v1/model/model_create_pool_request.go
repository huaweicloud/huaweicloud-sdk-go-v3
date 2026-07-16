package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePoolRequest Request Object
type CreatePoolRequest struct {

	// **参数解释**：实际的外部租户ID，如果有的话，工作空间鉴权以该ID为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	XModelArtsUserID *string `json:"X-ModelArts-User-ID,omitempty"`

	Body *CreatePoolRequestBody `json:"body,omitempty"`
}

func (o CreatePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolRequest struct{}"
	}

	return strings.Join([]string{"CreatePoolRequest", string(data)}, " ")
}
