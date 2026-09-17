package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterNodesInstallCmdRequest Request Object
type CreateClusterNodesInstallCmdRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`

	Body *CreateNodesInstallCmdV3RequestBody `json:"body,omitempty"`
}

func (o CreateClusterNodesInstallCmdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterNodesInstallCmdRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterNodesInstallCmdRequest", string(data)}, " ")
}
