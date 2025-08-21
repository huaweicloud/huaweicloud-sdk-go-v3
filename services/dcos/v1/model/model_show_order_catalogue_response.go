package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderCatalogueResponse Response Object
type ShowOrderCatalogueResponse struct {
	Body           *[]OrderCatalogue `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowOrderCatalogueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderCatalogueResponse struct{}"
	}

	return strings.Join([]string{"ShowOrderCatalogueResponse", string(data)}, " ")
}
