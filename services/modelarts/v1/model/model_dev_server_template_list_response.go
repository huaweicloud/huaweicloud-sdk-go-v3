package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DevServerTemplateListResponse DevServer模板列表
type DevServerTemplateListResponse struct {

	// **参数解释**：任务模板id。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：任务模板名。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：任务模板描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：任务模板类型。 **取值范围**：- COMMON  - LOG_COLLECT等。
	Type *string `json:"type,omitempty"`

	// **参数解释**：规格类型。 **取值范围**：-ASCEND_SNT9B   -ASCEND_SNT9C   -ASCEND_GENERIC。
	FlavorType *string `json:"flavor_type,omitempty"`

	// **参数解释**：模板的一些任务所需额外params参数。
	Params *[]TemplateParam `json:"params,omitempty"`
}

func (o DevServerTemplateListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevServerTemplateListResponse struct{}"
	}

	return strings.Join([]string{"DevServerTemplateListResponse", string(data)}, " ")
}
