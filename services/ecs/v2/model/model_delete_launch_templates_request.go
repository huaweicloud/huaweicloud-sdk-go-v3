package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLaunchTemplatesRequest Request Object
type DeleteLaunchTemplatesRequest struct {

	// 模板ID。
	LaunchTemplateId string `json:"launch_template_id"`
}

func (o DeleteLaunchTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLaunchTemplatesRequest struct{}"
	}

	return strings.Join([]string{"DeleteLaunchTemplatesRequest", string(data)}, " ")
}
