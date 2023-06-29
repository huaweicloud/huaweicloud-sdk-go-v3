package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectVpcChannelSyncsRequest Request Object
type CreateProjectVpcChannelSyncsRequest struct {
	Body *ProjectVpcSync `json:"body,omitempty"`
}

func (o CreateProjectVpcChannelSyncsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectVpcChannelSyncsRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectVpcChannelSyncsRequest", string(data)}, " ")
}
