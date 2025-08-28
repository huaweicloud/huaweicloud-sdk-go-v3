package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMemberRequest Request Object
type UpdateMemberRequest struct {

	// **参数解释**：后端服务器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及  > 此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	MemberId string `json:"member_id"`

	// **参数解释**：后端服务器组ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	PoolId string `json:"pool_id"`

	Body *UpdateMemberRequestBody `json:"body,omitempty"`
}

func (o UpdateMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberRequest", string(data)}, " ")
}
