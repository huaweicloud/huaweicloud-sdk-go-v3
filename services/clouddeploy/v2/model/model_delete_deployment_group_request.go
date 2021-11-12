package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeploymentGroupRequest struct {
	// 主机组id

	GroupId string `json:"group_id"`
}

func (o DeleteDeploymentGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentGroupRequest", string(data)}, " ")
}
