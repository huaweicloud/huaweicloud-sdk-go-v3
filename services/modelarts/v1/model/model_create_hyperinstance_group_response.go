package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHyperinstanceGroupResponse Response Object
type CreateHyperinstanceGroupResponse struct {

	// **参数解释**：hyperinstance-group ID。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：hyperinstance-group名称。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：npu数量限制。 **默认取值**：不涉及。
	NpuCountLimit *string `json:"npu_count_limit,omitempty"`

	// **参数解释**：npu数量。 **默认取值**：不涉及。
	NpuCount *string `json:"npu_count,omitempty"`

	// **参数解释**：超节点集群ID。 **默认取值**：不涉及。
	HyperinstanceClusterId *string `json:"hyperinstance_cluster_id,omitempty"`

	// **参数解释**：创建时间。 **默认取值**：不涉及。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// **参数解释**：更新时间。 **默认取值**：不涉及。
	UpdatedAt      *int64 `json:"updated_at,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateHyperinstanceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHyperinstanceGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateHyperinstanceGroupResponse", string(data)}, " ")
}
