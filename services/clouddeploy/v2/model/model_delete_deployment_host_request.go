package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeploymentHostRequest struct {

	// 主机组id
	GroupId string `json:"group_id"`

	// 主机id
	HostId string `json:"host_id"`
}

func (o DeleteDeploymentHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentHostRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentHostRequest", string(data)}, " ")
}
