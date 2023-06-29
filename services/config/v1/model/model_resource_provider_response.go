package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceProviderResponse 云服务详情
type ResourceProviderResponse struct {

	// 云服务名称
	Provider *string `json:"provider,omitempty"`

	// 云服务显示名称，可以通过请求Header中的'X-Language'设置语言
	DisplayName *string `json:"display_name,omitempty"`

	// 云服务类别显示名称，可以通过请求Header中的'X-Language'设置语言
	CategoryDisplayName *string `json:"category_display_name,omitempty"`

	// 资源类型列表
	ResourceTypes *[]ResourceTypeResponse `json:"resource_types,omitempty"`
}

func (o ResourceProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceProviderResponse struct{}"
	}

	return strings.Join([]string{"ResourceProviderResponse", string(data)}, " ")
}
