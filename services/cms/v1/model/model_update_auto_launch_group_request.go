package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoLaunchGroupRequest Request Object
type UpdateAutoLaunchGroupRequest struct {

	// 智能购买组ID
	AutoLaunchGroupId string `json:"auto_launch_group_id"`

	Body *UpdateAutoLaunchGroupReqV2 `json:"body,omitempty"`
}

func (o UpdateAutoLaunchGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoLaunchGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateAutoLaunchGroupRequest", string(data)}, " ")
}
