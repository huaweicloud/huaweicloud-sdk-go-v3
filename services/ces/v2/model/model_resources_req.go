package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesReq struct {

	// 资源信息
	Resources *[]Resource `json:"resources,omitempty"`

	// 手动创建，选择资源层级为云产品时的资源详情
	ProductResources *[]ProductResource `json:"product_resources,omitempty"`
}

func (o ResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesReq struct{}"
	}

	return strings.Join([]string{"ResourcesReq", string(data)}, " ")
}
