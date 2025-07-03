package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplatesRequest Request Object
type ListTemplatesRequest struct {

	// max number of resources to return
	Limit *int32 `json:"limit,omitempty"`

	// next page resource index id
	Marker *string `json:"marker,omitempty"`

	// 模板ID
	LaunchTemplateId *[]string `json:"launch_template_id,omitempty"`

	// 模板名称
	Name *[]string `json:"name,omitempty"`
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}
