package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchPoolRequest Request Object
type PatchPoolRequest struct {

	// **参数解释**：系统生成的资源池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	// **参数解释**：实际的外部租户ID，如果有的话，工作空间鉴权以该ID为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	XModelArtsUserID *string `json:"X-ModelArts-User-ID,omitempty"`

	Body *PoolUpdateRequest `json:"body,omitempty"`
}

func (o PatchPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchPoolRequest struct{}"
	}

	return strings.Join([]string{"PatchPoolRequest", string(data)}, " ")
}
