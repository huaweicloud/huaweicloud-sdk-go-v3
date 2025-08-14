package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateMetadataRequest Request Object
type UpdateTemplateMetadataRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 模板（Template）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
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
