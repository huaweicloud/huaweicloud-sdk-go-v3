package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLaunchTemplateResponse Response Object
type CreateLaunchTemplateResponse struct {

	// 模板id
	LaunchTemplateId *string `json:"launch_template_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLaunchTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLaunchTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateLaunchTemplateResponse", string(data)}, " ")
}
