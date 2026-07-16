package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDevServerJobTemplateResponse Response Object
type GetDevServerJobTemplateResponse struct {

	// **参数解释**：模板id。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：模板名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：模板描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：容器启动命令。 **取值范围**：不涉及。
	Cmd *string `json:"cmd,omitempty"`

	// **参数解释**：任务镜像。 **取值范围**：不涉及。
	SwrPath *string `json:"swr_path,omitempty"`

	// **参数解释**：任务规格。 **取值范围**：不涉及。
	Resources *string `json:"resources,omitempty"`

	// **参数解释**：卷。 **取值范围**：不涉及。
	Volumes *string `json:"volumes,omitempty"`

	// **参数解释**：卷挂载。 **取值范围**：不涉及。
	VolumesMount *string `json:"volumes_mount,omitempty"`

	// **参数解释**：规格类型。 **取值范围**：-ASCEND_SNT9B   -ASCEND_SNT9C   -ASCEND_GENERIC。
	FlavorType *string `json:"flavor_type,omitempty"`

	// **参数解释**：任务超时时间。 **取值范围**：不涉及。
	Timeout *string `json:"timeout,omitempty"`

	// **参数解释**：任务的轮询周期。 **取值范围**：不涉及。
	CheckInterval *string `json:"check_interval,omitempty"`

	// **参数解释**：任务类型。 **取值范围**：-LOG_COLLECT  -COMMON 等
	Type *string `json:"type,omitempty"`

	// **参数解释**：模板状态。 **取值范围**：ACTIVE。
	Status *string `json:"status,omitempty"`

	// **参数解释**：模板的其他参数。
	Params         *[]TemplateParam `json:"params,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o GetDevServerJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevServerJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"GetDevServerJobTemplateResponse", string(data)}, " ")
}
