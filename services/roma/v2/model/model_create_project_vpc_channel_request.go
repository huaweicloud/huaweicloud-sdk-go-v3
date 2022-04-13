package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProjectVpcChannelRequest struct {
	Body *ProjectVpcCreate `json:"body,omitempty"`
}

func (o CreateProjectVpcChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectVpcChannelRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectVpcChannelRequest", string(data)}, " ")
}
