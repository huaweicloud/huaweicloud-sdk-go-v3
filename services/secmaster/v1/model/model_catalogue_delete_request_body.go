package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogueDeleteRequestBody 删除布局请求体
type CatalogueDeleteRequestBody struct {

	// 布局ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`

	// 是否删除，为true时直接删除；如有引用关系，则删除失败
	IsDelete *bool `json:"is_delete,omitempty"`
}

func (o CatalogueDeleteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogueDeleteRequestBody struct{}"
	}

	return strings.Join([]string{"CatalogueDeleteRequestBody", string(data)}, " ")
}
