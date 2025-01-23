package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceNodesInfoResponse Response Object
type ListInstanceNodesInfoResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例节点列表
	InstanceNodes  *[]InstanceNodesInfoInstanceNodes `json:"instance_nodes,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListInstanceNodesInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceNodesInfoResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceNodesInfoResponse", string(data)}, " ")
}
