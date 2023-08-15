package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CatalogParamsVo struct {
	Entity *CatalogEntityVo `json:"entity"`
}

func (o CatalogParamsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogParamsVo struct{}"
	}

	return strings.Join([]string{"CatalogParamsVo", string(data)}, " ")
}
