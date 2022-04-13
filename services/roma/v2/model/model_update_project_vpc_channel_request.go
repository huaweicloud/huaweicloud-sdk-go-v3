package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProjectVpcChannelRequest struct {
	Body *ProjectVpcCreate `json:"body,omitempty"`
}

func (o UpdateProjectVpcChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectVpcChannelRequest struct{}"
	}

	return strings.Join([]string{"UpdateProjectVpcChannelRequest", string(data)}, " ")
}
