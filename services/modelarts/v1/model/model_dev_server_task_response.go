package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DevServerTaskResponse DevServerTask详情
type DevServerTaskResponse struct {

	// **参数解释**：task的ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：devserver机器ID。 **取值范围**：不涉及。
	ServerId *string `json:"server_id,omitempty"`

	// **参数解释**：devserver机器名称。 **取值范围**：不涉及。
	ServerName *string `json:"server_name,omitempty"`

	// **参数解释**：task状态。 **取值范围**：- PROCESSING  -SUCCESS  - FAILED  - SKIPPED
	Status *string `json:"status,omitempty"`

	// **参数解释**：底层ECS/BMS/HPS ID。
	CloudServer map[string]string `json:"cloud_server,omitempty"`

	// **参数解释**：输出信息。 **取值范围**：不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释**：更新时间。 **取值范围**：不涉及。
	UpdateAt *string `json:"update_at,omitempty"`
}

func (o DevServerTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevServerTaskResponse struct{}"
	}

	return strings.Join([]string{"DevServerTaskResponse", string(data)}, " ")
}
