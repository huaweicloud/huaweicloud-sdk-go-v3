package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeploymentGroupResponse Response Object
type UpdateDeploymentGroupResponse struct {

	// 主机集群ID
	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDeploymentGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeploymentGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeploymentGroupResponse", string(data)}, " ")
}
