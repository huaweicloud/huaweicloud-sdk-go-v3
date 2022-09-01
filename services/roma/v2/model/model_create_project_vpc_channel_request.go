package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProjectVpcChannelRequest struct {
	Body *ProjectVpcCreate `json:"body,omitempty" xml:"body"`
}

func (o CreateProjectVpcChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectVpcChannelRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectVpcChannelRequest", string(data)}, " ")
}
