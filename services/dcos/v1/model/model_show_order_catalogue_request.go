package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderCatalogueRequest Request Object
type ShowOrderCatalogueRequest struct {
}

func (o ShowOrderCatalogueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderCatalogueRequest struct{}"
	}

	return strings.Join([]string{"ShowOrderCatalogueRequest", string(data)}, " ")
}
