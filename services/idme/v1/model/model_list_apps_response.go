package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppsResponse Response Object
type ListAppsResponse struct {

	// 分页查询时，每页最多展示的记录数。
	PageNum *int32 `json:"page_num,omitempty"`

	// 分页查询的页数。
	PageSize *int32 `json:"page_size,omitempty"`

	// 总共条数。
	TotalSize *int32 `json:"total_size,omitempty"`

	// 总页数。
	TotalPages *int32 `json:"total_pages,omitempty"`

	// 应用详情。
	Result         *[]ListAppsResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsResponse struct{}"
	}

	return strings.Join([]string{"ListAppsResponse", string(data)}, " ")
}
