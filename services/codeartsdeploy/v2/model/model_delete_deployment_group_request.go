package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentGroupRequest Request Object
type DeleteDeploymentGroupRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`
}

func (o DeleteDeploymentGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentGroupRequest", string(data)}, " ")
}
