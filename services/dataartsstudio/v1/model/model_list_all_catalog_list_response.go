package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllCatalogListResponse Response Object
type ListAllCatalogListResponse struct {

	// 符合条件的数据总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回的APP列表
	ApiCatalogs    *[]RecordForGetAllCatalog `json:"api_catalogs,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAllCatalogListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllCatalogListResponse struct{}"
	}

	return strings.Join([]string{"ListAllCatalogListResponse", string(data)}, " ")
}
