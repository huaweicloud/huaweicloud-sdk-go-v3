package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProjectVpcChannelRequest struct {
	Body *ProjectVpcCreate `json:"body,omitempty" xml:"body"`
}

func (o UpdateProjectVpcChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectVpcChannelRequest struct{}"
	}

	return strings.Join([]string{"UpdateProjectVpcChannelRequest", string(data)}, " ")
}
