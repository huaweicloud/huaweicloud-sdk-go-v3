package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDevServerJobResponse Response Object
type GetDevServerJobResponse struct {

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释**：更新时间。 **取值范围**：不涉及。
	UpdateAt *string `json:"update_at,omitempty"`

	// **参数解释**：任务ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：任务名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：任务模板类型。 **取值范围**：- COMMON  -SERVICE_DEPLOY 等。
	Type *string `json:"type,omitempty"`

	// **参数解释**：状态。 **取值范围**：- ACTIVE。
	Status *string `json:"status,omitempty"`

	// **参数解释**：任务实例列表信息。
	Items *[]DevServerJobItem `json:"items,omitempty"`

	// **参数解释**：task详情列表。
	Tasks *[]DevServerTaskResponse `json:"tasks,omitempty"`

	// **参数解释**：任务模板ID。 **取值范围**：不涉及。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释**：下发任务的用户信息。 **取值范围**：不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：task失败的节点数量。 **取值范围**：不涉及。
	AbnormalCount *int32 `json:"abnormal_count,omitempty"`

	// **参数解释**：描述。 **取值范围**：不涉及。
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetDevServerJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevServerJobResponse struct{}"
	}

	return strings.Join([]string{"GetDevServerJobResponse", string(data)}, " ")
}
