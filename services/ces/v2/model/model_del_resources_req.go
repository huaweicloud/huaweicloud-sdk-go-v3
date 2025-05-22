package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DelResourcesReq struct {

	// 手动创建，选择资源层级为子维度时的资源信息。只需要传递删除的资源
	Resources *[]Resource `json:"resources,omitempty"`

	// 手动创建，选择资源层级为云产品时的资源详情。只需要传递删除的资源
	ProductResources *[]ProductResource `json:"product_resources,omitempty"`
}

func (o DelResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelResourcesReq struct{}"
	}

	return strings.Join([]string{"DelResourcesReq", string(data)}, " ")
}
