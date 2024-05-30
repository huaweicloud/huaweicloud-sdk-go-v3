package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCatalogResponse Response Object
type ChangeCatalogResponse struct {
	Data           *CreateCatalogResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ChangeCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCatalogResponse struct{}"
	}

	return strings.Join([]string{"ChangeCatalogResponse", string(data)}, " ")
}
