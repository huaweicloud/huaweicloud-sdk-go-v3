package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodesInformationResponse Response Object
type ShowNodesInformationResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 节点数量
	Total *int32 `json:"total,omitempty"`

	// 节点信息
	Nodes          *[]NodesInfoResp `json:"nodes,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowNodesInformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodesInformationResponse struct{}"
	}

	return strings.Join([]string{"ShowNodesInformationResponse", string(data)}, " ")
}
