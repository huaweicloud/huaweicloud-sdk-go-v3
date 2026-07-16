package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// **参数解释**：任务类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - replace-node：故障节点替换任务。 - reboot-node：节点重启任务。 - reset-nodes：节点重置任务。 **默认取值**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：单页查询最大数量，该值为空时默认返回100条记录，单页最大允许查询500条记录。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页查询的偏移量，查询第一页时为空。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：查询起始时间，单位毫秒。默认从30天前开始查询。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Since *int64 `json:"since,omitempty"`

	// **参数解释**：查询终止时间，单位毫秒。默认当前时间。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Until *int64 `json:"until,omitempty"`

	// **参数解释**：任务关联的资源类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - pools：资源池 **默认取值**：不涉及。
	Resource *string `json:"resource,omitempty"`

	// **参数解释**：关联的资源名称，与resource配合使用。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
