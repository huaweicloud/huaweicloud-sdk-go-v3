package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLaunchTemplateRequestBody struct {
	LaunchTemplate *LaunchTemplateOption `json:"launch_template"`

	// 预检，只校验权限等初级信息。
	DryRun *bool `json:"dry_run,omitempty"`
}

func (o CreateLaunchTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLaunchTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLaunchTemplateRequestBody", string(data)}, " ")
}
