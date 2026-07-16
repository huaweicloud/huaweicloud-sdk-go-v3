package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodePoolsRequest Request Object
type ListNodePoolsRequest struct {

	// **参数解释**：资源池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`
}

func (o ListNodePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListNodePoolsRequest", string(data)}, " ")
}
