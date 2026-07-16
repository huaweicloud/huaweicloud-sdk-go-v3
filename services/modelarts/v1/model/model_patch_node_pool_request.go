package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchNodePoolRequest Request Object
type PatchNodePoolRequest struct {

	// **参数解释**：资源池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	// **参数解释**：节点池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodepoolName string `json:"nodepool_name"`

	Body *PatchNodePoolRequestBody `json:"body,omitempty"`
}

func (o PatchNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchNodePoolRequest struct{}"
	}

	return strings.Join([]string{"PatchNodePoolRequest", string(data)}, " ")
}
