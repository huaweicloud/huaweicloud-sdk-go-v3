package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperClusterRequest Request Object
type ListHyperClusterRequest struct {

	// **参数解释**：Hyper Cluster的类型。 **约束限制**：可选。 **取值范围**： - HPS：默认值，查询HPS机型的Hyper Cluster。 - ECS：查询ECS机型的Hyper Cluster。  **默认取值**：HPS。
	Type *string `json:"type,omitempty"`
}

func (o ListHyperClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperClusterRequest struct{}"
	}

	return strings.Join([]string{"ListHyperClusterRequest", string(data)}, " ")
}
