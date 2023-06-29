package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LayerPath struct {

	// 目录编号
	CatalogId *string `json:"catalog_id,omitempty"`

	// 路径层级
	Order *int32 `json:"order,omitempty"`
}

func (o LayerPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayerPath struct{}"
	}

	return strings.Join([]string{"LayerPath", string(data)}, " ")
}
