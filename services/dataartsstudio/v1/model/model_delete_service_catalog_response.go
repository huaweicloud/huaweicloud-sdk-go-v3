package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceCatalogResponse Response Object
type DeleteServiceCatalogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServiceCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceCatalogResponse struct{}"
	}

	return strings.Join([]string{"DeleteServiceCatalogResponse", string(data)}, " ")
}
