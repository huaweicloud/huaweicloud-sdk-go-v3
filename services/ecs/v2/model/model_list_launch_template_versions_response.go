package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLaunchTemplateVersionsResponse Response Object
type ListLaunchTemplateVersionsResponse struct {

	// 模板版本列表
	LaunchTemplateVersions *[]TemplateVersion `json:"launch_template_versions,omitempty"`

	PageInfo *ResponsePageInfoV3 `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLaunchTemplateVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLaunchTemplateVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListLaunchTemplateVersionsResponse", string(data)}, " ")
}
