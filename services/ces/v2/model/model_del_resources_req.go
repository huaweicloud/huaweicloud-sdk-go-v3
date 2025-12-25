package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DelResourcesReq struct {

	// **参数解释** 当资源添加方式为手动创建、资源层级为子维度时，删除资源分组中的资源只需传递删除的资源信息 **约束限制** 包含的资源对象个数为1到1000
	Resources *[]Resource `json:"resources,omitempty"`

	// **参数解释** 当资源添加方式为手动创建、资源层级为子维度时，删除资源分组中的资源只需传递删除的资源信息 **约束限制** 包含的资源个数为1到50个
	ProductResources *[]ProductResource `json:"product_resources,omitempty"`
}

func (o DelResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelResourcesReq struct{}"
	}

	return strings.Join([]string{"DelResourcesReq", string(data)}, " ")
}
