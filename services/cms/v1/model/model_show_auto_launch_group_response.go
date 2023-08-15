package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoLaunchGroupResponse Response Object
type ShowAutoLaunchGroupResponse struct {
	AutoLaunchGroup *ShowAutoLaunchGroupResp `json:"auto_launch_group,omitempty"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ShowAutoLaunchGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoLaunchGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoLaunchGroupResponse", string(data)}, " ")
}
