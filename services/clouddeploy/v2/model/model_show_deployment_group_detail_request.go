package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeploymentGroupDetailRequest struct {
	// 主机组id

	GroupId string `json:"group_id"`
}

func (o ShowDeploymentGroupDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentGroupDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentGroupDetailRequest", string(data)}, " ")
}
