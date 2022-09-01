package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源类型详情
type ResourceTypeResponse struct {

	// 资源类型名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 资源类型显示名称，可以通过请求中 'X-Language'设置语言
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 是否是全局类型的资源
	Global *bool `json:"global,omitempty" xml:"global"`

	// 支持的region列表
	Regions *[]string `json:"regions,omitempty" xml:"regions"`

	// console终端id
	ConsoleEndpointId *string `json:"console_endpoint_id,omitempty" xml:"console_endpoint_id"`

	// console列表页地址
	ConsoleListUrl *string `json:"console_list_url,omitempty" xml:"console_list_url"`

	// console详情页地址
	ConsoleDetailUrl *string `json:"console_detail_url,omitempty" xml:"console_detail_url"`
}

func (o ResourceTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTypeResponse struct{}"
	}

	return strings.Join([]string{"ResourceTypeResponse", string(data)}, " ")
}
