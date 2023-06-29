package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogInput catalog信息
type CatalogInput struct {

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 路径地址
	Location string `json:"location"`

	// database路径列表
	DatabaseLocationList *[]string `json:"database_location_list,omitempty"`
}

func (o CatalogInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogInput struct{}"
	}

	return strings.Join([]string{"CatalogInput", string(data)}, " ")
}
