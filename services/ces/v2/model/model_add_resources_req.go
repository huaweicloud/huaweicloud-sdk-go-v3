package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddResourcesReq struct {

	// 手动创建，选择资源层级为子维度时的资源信息。每次修改资源时只需要传递新增的资源
	Resources *[]Resource `json:"resources,omitempty"`

	// 手动创建，选择资源层级为云产品时的资源详情。每次修改资源时需要将已有资源和新增的资源一起传递
	ProductResources *[]ProductResource `json:"product_resources,omitempty"`
}

func (o AddResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResourcesReq struct{}"
	}

	return strings.Join([]string{"AddResourcesReq", string(data)}, " ")
}
