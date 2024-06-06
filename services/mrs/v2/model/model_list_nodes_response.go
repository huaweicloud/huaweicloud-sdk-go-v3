package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodesResponse Response Object
type ListNodesResponse struct {

	// 节点列表。
	Nodes *[]ClusterNode `json:"nodes,omitempty"`

	// 节点数。
	NodeTotal      *int32 `json:"node_total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNodesResponse", string(data)}, " ")
}
