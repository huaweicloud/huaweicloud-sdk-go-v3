package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteMembersState struct {

	// **参数解释**：后端服务器ID。  **取值范围**：不涉及  > 此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	Id string `json:"id"`

	// **参数解释**：当前后端服务器删除结果状态。  **取值范围**： - successful：删除成功。 - not found：member不存在。
	RetStatus string `json:"ret_status"`
}

func (o BatchDeleteMembersState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersState struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersState", string(data)}, " ")
}
