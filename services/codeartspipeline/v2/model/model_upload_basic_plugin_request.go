package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadBasicPluginRequest Request Object
type UploadBasicPluginRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 插件名
	PluginName string `json:"plugin_name"`

	// 业务类型
	BusinessType string `json:"business_type"`

	Body *UploadBasicPluginRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadBasicPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadBasicPluginRequest struct{}"
	}

	return strings.Join([]string{"UploadBasicPluginRequest", string(data)}, " ")
}
