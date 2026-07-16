package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodePoolRequest Request Object
type ShowNodePoolRequest struct {

	// **参数解释**：资源池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	// **参数解释**：节点池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodepoolName string `json:"nodepool_name"`

	// **参数解释**：分页查询时上一页位置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Continue *string `json:"continue,omitempty"`

	// **参数解释**：分页单次查询返回数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Limit *string `json:"limit,omitempty"`
}

func (o ShowNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodePoolRequest struct{}"
	}

	return strings.Join([]string{"ShowNodePoolRequest", string(data)}, " ")
}
