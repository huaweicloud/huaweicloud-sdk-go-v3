package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyHostsToTargetRequest Request Object
type CopyHostsToTargetRequest struct {

	// 源主机集群id
	GroupId string `json:"group_id"`

	Body *DeploymentHostsCopyRequest `json:"body,omitempty"`
}

func (o CopyHostsToTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyHostsToTargetRequest struct{}"
	}

	return strings.Join([]string{"CopyHostsToTargetRequest", string(data)}, " ")
}
