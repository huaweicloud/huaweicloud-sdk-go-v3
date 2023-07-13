package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeploymentGroupDetailRequest Request Object
type ShowDeploymentGroupDetailRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`
}

func (o ShowDeploymentGroupDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentGroupDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentGroupDetailRequest", string(data)}, " ")
}
