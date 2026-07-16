package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DevServerJobCreateRequest struct {

	// **参数解释**：任务名称。 **约束限制**：^[-_.a-zA-Z0-9]{1,64}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：任务描述。 **约束限制**：^[-_.a-zA-Z0-9]{1,64}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：DevServer实例id列表。 **取值范围**：不涉及。
	ServerIds []string `json:"server_ids"`

	// **参数解释**：任务模板类型。 **约束限制**：^[-_.a-zA-Z0-9]{1,64}$。 **取值范围**：-COMMON  -SERVICE_DEPLOY 等。 **默认取值**：不涉及。
	Type string `json:"type"`

	// **参数解释**：任务失败后是否重启。 **约束限制**：不涉及。 **取值范围**：- true   -false。 **默认取值**：false。
	IsReboot *bool `json:"is_reboot,omitempty"`

	// **参数解释**：任务实例列表。 **取值范围**：不涉及。
	Items []DevServerJobItem `json:"items"`
}

func (o DevServerJobCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevServerJobCreateRequest struct{}"
	}

	return strings.Join([]string{"DevServerJobCreateRequest", string(data)}, " ")
}
