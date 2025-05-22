package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterCnResponse Response Object
type ListClusterCnResponse struct {

	// **参数解释**： 集群支持的最小CN节点数量。 **取值范围**： 不涉及。
	MinNum *int32 `json:"min_num,omitempty"`

	// **参数解释**： 集群支持的最大CN节点数量。 **取值范围**： 不涉及。
	MaxNum *int32 `json:"max_num,omitempty"`

	// **参数解释**： CN节点详情列表。 **取值范围**： 不涉及。
	Instances      *[]CoordinatorNode `json:"instances,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListClusterCnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterCnResponse struct{}"
	}

	return strings.Join([]string{"ListClusterCnResponse", string(data)}, " ")
}
