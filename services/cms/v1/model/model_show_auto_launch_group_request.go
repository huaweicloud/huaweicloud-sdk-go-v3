package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAutoLaunchGroupRequest struct {

	// 智能购买组id
	AutoLaunchGroupId string `json:"auto_launch_group_id"`
}

func (o ShowAutoLaunchGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoLaunchGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoLaunchGroupRequest", string(data)}, " ")
}
