package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddResourcesReq struct {

	// 当资源添加方式为手动创建、资源层级为子维度时，资源分组新增资源时只需传递新增的资源信息
	Resources *[]Resource `json:"resources,omitempty"`

	// 当资源添加方式为手动创建、资源层级为云产品时，资源分组新增资源时需要将已有资源信息和新增的资源信息一起传递
	ProductResources *[]ProductResource `json:"product_resources,omitempty"`
}

func (o AddResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResourcesReq struct{}"
	}

	return strings.Join([]string{"AddResourcesReq", string(data)}, " ")
}
