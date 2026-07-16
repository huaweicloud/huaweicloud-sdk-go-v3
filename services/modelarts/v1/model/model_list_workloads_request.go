package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadsRequest Request Object
type ListWorkloadsRequest struct {

	// **参数解释**：资源池的ID，取值自资源池详情的metadata.name字段。 **约束限制**：不涉及。 **取值范围**：只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为[36-63]个字符。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	// **参数解释**：节点IP地址，用于过滤在该节点上运行的workload。支持多个IP，传递多个参数或用逗号分隔. **约束限制**：指定此参数时，不支持分页（limit/offset会被忽略）。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Hostip *[]string `json:"hostip,omitempty"`

	// **参数解释**：根据作业类型查询资源池作业列表。 **约束限制**：不涉及。 **取值范围**：可选值如下： - train：训练作业 - infer：推理服务 - notebook：Notebook作业 - x-infer：新版推理作业 **默认取值**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：根据作业状态查询资源池作业列表。 **约束限制**：不涉及。 **取值范围**：可选值如下： - Queue：排队中的作业。 - Pending：等待中的作业。 - Abnormal：异常的作业。 - Terminating：中止中的作业。 - Creating：创建中的作业。 - Running：运行中的作业。 - Completed：已完成的作业。 - Terminated：已终止的作业。 - Failed：运行失败的作业。 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：查询资源池作业列表的排序条件。 **约束限制**：不涉及。 **取值范围**：可选值如下： - create_time：根据作业创建时间排序。 **默认取值**：不涉及。
	Sort *string `json:"sort,omitempty"`

	// **参数解释**：指定查询资源池作业列表是否按照升序排序。 **约束限制**：需要配合sort查询参数使用。 **取值范围**：可选值如下： - true：按照升序排序。 - false：按照降序排序。 **默认取值**：false。
	Ascend *bool `json:"ascend,omitempty"`

	// **参数解释**：分页查询的偏移量。 **约束限制**：不涉及。 **取值范围**：0-2147483647。 **默认取值**：0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释**：分页单次查询返回的资源数量。 **约束限制**：不涉及。 **取值范围**：0 - 500。 **默认取值**：500。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListWorkloadsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkloadsRequest", string(data)}, " ")
}
