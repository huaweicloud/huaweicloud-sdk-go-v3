package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateMetadataRequest Request Object
type UpdateTemplateMetadataRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 用户希望创建的模板名称
	TemplateName string `json:"template_name"`

	Body *UpdateTemplateMetadataRequestBody `json:"body,omitempty"`
}

func (o UpdateTemplateMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateMetadataRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateMetadataRequest", string(data)}, " ")
}
