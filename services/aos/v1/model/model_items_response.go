package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// items response
type ItemsResponse struct {

	// 执行计划中的资源类型，如：huaweicloud_evs_volume
	ResourceType *string `json:"resource_type,omitempty"`

	// 执行计划中的用户定义的资源的名字，如：my_volume
	ResourceName *string `json:"resource_name,omitempty"`

	// 使用count构建的资源时资源对应的index
	Index *string `json:"index,omitempty"`

	// 执行计划中的资源是否支持询价
	Supported *bool `json:"supported,omitempty"`

	// 执行计划中的每个资源部署后的总的询价信息，如果该资源询价结果包含不同的period_type，则询价结果根据period_type类型展示总价。 包周期计费和按需计费返回，免费和不支持询价的资源不返回
	ResourcePrice *[]ResourcePriceResponse `json:"resource_price,omitempty"`
}

func (o ItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemsResponse struct{}"
	}

	return strings.Join([]string{"ItemsResponse", string(data)}, " ")
}
