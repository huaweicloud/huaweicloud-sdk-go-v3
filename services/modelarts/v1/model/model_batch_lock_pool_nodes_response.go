package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchLockPoolNodesResponse Response Object
type BatchLockPoolNodesResponse struct {

	// **参数解释**：批量操作的资源池节点ID列表及状态。
	Nodes          *[]NodesResultMsg `json:"nodes,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchLockPoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchLockPoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchLockPoolNodesResponse", string(data)}, " ")
}
