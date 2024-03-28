package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogsResponse Response Object
type ListCatalogsResponse struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// catalog总数量
	TotalCount *int64 `json:"total_count,omitempty"`

	// 项目下所有catalog信息
	Catalogs       *[]Catalog `json:"catalogs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListCatalogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogsResponse struct{}"
	}

	return strings.Join([]string{"ListCatalogsResponse", string(data)}, " ")
}
