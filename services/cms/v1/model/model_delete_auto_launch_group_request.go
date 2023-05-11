package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAutoLaunchGroupRequest struct {

	// 智能购买组id
	AutoLaunchGroupId string `json:"auto_launch_group_id"`

	Body *DeleteAutoLaunchGroupReqV2 `json:"body,omitempty"`
}

func (o DeleteAutoLaunchGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutoLaunchGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteAutoLaunchGroupRequest", string(data)}, " ")
}
