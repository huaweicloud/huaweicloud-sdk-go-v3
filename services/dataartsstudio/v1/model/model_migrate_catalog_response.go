package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateCatalogResponse Response Object
type MigrateCatalogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MigrateCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateCatalogResponse struct{}"
	}

	return strings.Join([]string{"MigrateCatalogResponse", string(data)}, " ")
}
