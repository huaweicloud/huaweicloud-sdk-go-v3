package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// catalog response
type Catalog struct {

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 路径地址
	Location string `json:"location"`

	// database路径列表
	DatabaseLocationList *[]string `json:"database_location_list,omitempty"`
}

func (o Catalog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Catalog struct{}"
	}

	return strings.Join([]string{"Catalog", string(data)}, " ")
}
