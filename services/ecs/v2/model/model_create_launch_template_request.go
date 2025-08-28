package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLaunchTemplateRequest Request Object
type CreateLaunchTemplateRequest struct {
	Body *CreateLaunchTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateLaunchTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLaunchTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateLaunchTemplateRequest", string(data)}, " ")
}
