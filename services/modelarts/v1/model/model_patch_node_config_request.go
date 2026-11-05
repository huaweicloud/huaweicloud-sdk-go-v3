package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchNodeConfigRequest Request Object
type PatchNodeConfigRequest struct {

	// **参数解释**：池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	Body *NodeconfigUpdate `json:"body,omitempty"`
}

func (o PatchNodeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchNodeConfigRequest struct{}"
	}

	return strings.Join([]string{"PatchNodeConfigRequest", string(data)}, " ")
}
