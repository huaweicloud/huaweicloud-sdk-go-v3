package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCatalogRequest Request Object
type ShowCatalogRequest struct {

	// catalog名称。
	CatalogName string `json:"catalog_name"`
}

func (o ShowCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCatalogRequest struct{}"
	}

	return strings.Join([]string{"ShowCatalogRequest", string(data)}, " ")
}
