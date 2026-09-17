package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterNodesResponse Response Object
type ListClusterNodesResponse struct {

	// 节点列表
	Nodes          *[]QueryNodeResp `json:"nodes,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListClusterNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterNodesResponse struct{}"
	}

	return strings.Join([]string{"ListClusterNodesResponse", string(data)}, " ")
}
