package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiCatalogListResponse Response Object
type ListApiCatalogListResponse struct {

	// API总数量
	Total *int32 `json:"total,omitempty"`

	// API列表
	Apis           *[]ApiOverview `json:"apis,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListApiCatalogListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiCatalogListResponse struct{}"
	}

	return strings.Join([]string{"ListApiCatalogListResponse", string(data)}, " ")
}
