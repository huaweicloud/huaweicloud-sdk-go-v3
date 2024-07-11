package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostInfoRequest Request Object
type UpdateHostInfoRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	// 主机id
	HostId string `json:"host_id"`

	Body *DeploymentHostRequestExternal `json:"body,omitempty"`
}

func (o UpdateHostInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostInfoRequest", string(data)}, " ")
}
