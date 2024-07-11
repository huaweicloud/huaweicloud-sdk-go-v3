package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentGroupResponse Response Object
type DeleteDeploymentGroupResponse struct {

	// 主机集群id
	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeploymentGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentGroupResponse", string(data)}, " ")
}
