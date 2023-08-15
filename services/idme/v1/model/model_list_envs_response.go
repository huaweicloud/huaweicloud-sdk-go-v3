package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvsResponse Response Object
type ListEnvsResponse struct {

	// 分页查询的页数。
	PageNum *int32 `json:"page_num,omitempty"`

	// 分页查询时，每页最多展示的记录数。
	PageSize *int32 `json:"page_size,omitempty"`

	// 总共条数。
	TotalSize *int32 `json:"total_size,omitempty"`

	// 总共页数。
	TotalPages *int32 `json:"total_pages,omitempty"`

	// 运行服务详情。
	Result         *[]ListEnvsResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListEnvsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvsResponse", string(data)}, " ")
}
