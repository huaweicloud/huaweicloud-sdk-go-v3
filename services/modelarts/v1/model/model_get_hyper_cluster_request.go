package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetHyperClusterRequest Request Object
type GetHyperClusterRequest struct {

	// **参数解释**：Hyper Cluster ID。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id string `json:"id"`

	// **参数解释**：Hyper Cluster的类型。 **约束限制**：可选。 **取值范围**： - HPS：默认值，查询HPS机型的Hyper Cluster。 - ECS：查询ECS机型的Hyper Cluster。  **默认取值**：HPS。
	Type *string `json:"type,omitempty"`
}

func (o GetHyperClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHyperClusterRequest struct{}"
	}

	return strings.Join([]string{"GetHyperClusterRequest", string(data)}, " ")
}
