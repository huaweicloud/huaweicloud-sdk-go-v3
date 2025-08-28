package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteMembersOption struct {

	// **参数解释**：需要删除的后端服务器ID。  **约束限制**： - 若传入id则不能传其他参数，否则报错。  **取值范围**：不涉及  **默认取值**：不涉及  > 此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	Id *string `json:"id,omitempty"`

	// **参数解释**：后端服务器IP地址。  **约束限制**： - address和protocol_port必须同时传入。 - 不能同时传入ID字段。  **取值范围**：不涉及  **默认取值**：不涉及
	Address *string `json:"address,omitempty"`

	// **参数解释**：后端服务器端口。  **约束限制**： - address和protocol_port必须同时传入。 - 不能同时传入ID字段。 - 可以传0，用于删除端口透传pool下的member。  **取值范围**：不涉及  **默认取值**：不涉及
	ProtocolPort *int32 `json:"protocol_port,omitempty"`
}

func (o BatchDeleteMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersOption", string(data)}, " ")
}
