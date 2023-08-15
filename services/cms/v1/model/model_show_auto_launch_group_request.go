package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoLaunchGroupRequest Request Object
type ShowAutoLaunchGroupRequest struct {

	// 智能购买组ID
	AutoLaunchGroupId string `json:"auto_launch_group_id"`
}

func (o ShowAutoLaunchGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoLaunchGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoLaunchGroupRequest", string(data)}, " ")
}
