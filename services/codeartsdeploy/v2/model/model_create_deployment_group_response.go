package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeploymentGroupResponse Response Object
type CreateDeploymentGroupResponse struct {

	// 主机集群id
	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeploymentGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateDeploymentGroupResponse", string(data)}, " ")
}
