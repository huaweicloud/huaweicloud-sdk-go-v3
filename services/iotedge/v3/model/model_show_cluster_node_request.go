package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterNodeRequest Request Object
type ShowClusterNodeRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`

	// 节点名称
	NodeName string `json:"node_name"`
}

func (o ShowClusterNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterNodeRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterNodeRequest", string(data)}, " ")
}
