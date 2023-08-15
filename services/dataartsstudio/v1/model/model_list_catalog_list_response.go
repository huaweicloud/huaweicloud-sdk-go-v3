package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogListResponse Response Object
type ListCatalogListResponse struct {

	// 符合条件的数据总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回的APP列表
	Catalogs       *[]RecordForGetAllCatalog `json:"catalogs,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListCatalogListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogListResponse struct{}"
	}

	return strings.Join([]string{"ListCatalogListResponse", string(data)}, " ")
}
