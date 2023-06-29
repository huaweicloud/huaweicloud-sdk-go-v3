package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterInstallCmdRequest Request Object
type CreateClusterInstallCmdRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`

	// 边缘集群架构
	Arch string `json:"arch"`

	// 集群操作系统内核
	Os *string `json:"os,omitempty"`
}

func (o CreateClusterInstallCmdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInstallCmdRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterInstallCmdRequest", string(data)}, " ")
}
