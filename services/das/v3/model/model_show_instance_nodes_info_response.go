package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceNodesInfoResponse Response Object
type ShowInstanceNodesInfoResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例节点列表
	Nodes          *[]NodeInfo `json:"nodes,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowInstanceNodesInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceNodesInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceNodesInfoResponse", string(data)}, " ")
}
