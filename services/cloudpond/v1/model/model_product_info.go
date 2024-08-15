package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductInfo 商品归属产品属性信息
type ProductInfo struct {

	// 云服务类型
	CloudServiceType string `json:"cloud_service_type"`

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 资源编码
	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}
