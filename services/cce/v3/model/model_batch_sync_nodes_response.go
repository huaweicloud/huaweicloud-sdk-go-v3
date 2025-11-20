package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSyncNodesResponse Response Object
type BatchSyncNodesResponse struct {

	// **参数解释**： 固定值\"Start to batch sync nodes\"，表示批量同步节点成功。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchSyncNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSyncNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchSyncNodesResponse", string(data)}, " ")
}
