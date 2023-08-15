package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CatalogEntityVo struct {

	// 类型名称,填写“BusinessCatalog”即可
	TypeName *string `json:"typeName,omitempty"`

	Attributes *CatalogAttributeVo `json:"attributes"`
}

func (o CatalogEntityVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogEntityVo struct{}"
	}

	return strings.Join([]string{"CatalogEntityVo", string(data)}, " ")
}
