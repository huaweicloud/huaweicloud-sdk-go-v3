package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DevServerJobListResponse DevServerJob任务列表
type DevServerJobListResponse struct {

	// **参数解释**：任务id。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：任务名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：用户名。 **取值范围**：不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：任务描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：任务类型。 **取值范围**：- COMMON  -SERVICE_DEPLOY  等。
	Type *string `json:"type,omitempty"`

	// **参数解释**：任务状态。 **取值范围**：- PROCESSING,  -FINISHED,  -DELETED。
	Status *string `json:"status,omitempty"`

	// **参数解释**：task失败的节点数量。 **取值范围**：不涉及。
	AbnormalCount *int32 `json:"abnormal_count,omitempty"`

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释**：更新时间。 **取值范围**：不涉及。
	UpdateAt *string `json:"update_at,omitempty"`
}

func (o DevServerJobListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevServerJobListResponse struct{}"
	}

	return strings.Join([]string{"DevServerJobListResponse", string(data)}, " ")
}
