package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoViewDto struct {

	// 当前页。
	CurPage *int32 `json:"curPage,omitempty"`

	// 每页大小。
	PageSize *int32 `json:"pageSize,omitempty"`

	// 总行数。
	TotalRows *int32 `json:"totalRows,omitempty"`

	// 总页数。
	TotalPages *int32 `json:"totalPages,omitempty"`
}

func (o PageInfoViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoViewDto struct{}"
	}

	return strings.Join([]string{"PageInfoViewDto", string(data)}, " ")
}
