package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Catalog 分类信息描述。
type Catalog struct {

	// 唯一标识ID。
	Id *string `json:"id,omitempty"`

	// 分类描述(中文)。
	CatalogZh *string `json:"catalog_zh,omitempty"`

	// 分类描述(英文)。
	CatalogEn *string `json:"catalog_en,omitempty"`
}

func (o Catalog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Catalog struct{}"
	}

	return strings.Join([]string{"Catalog", string(data)}, " ")
}
