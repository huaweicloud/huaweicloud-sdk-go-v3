package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGetChildListRequest Request Object
type ListGetChildListRequest struct {

	// 当前页。
	CurPage *int32 `json:"curPage,omitempty"`

	// 结束索引。
	EndIndex *int32 `json:"endIndex,omitempty"`

	// 最大分页数。
	MaxPageSize *int32 `json:"maxPageSize,omitempty"`

	// 每页大小。
	PageSize *int32 `json:"pageSize,omitempty"`

	// 起始索引。
	StartIndex *int32 `json:"startIndex,omitempty"`

	// 总页数。
	TotalPages *int32 `json:"totalPages,omitempty"`

	// 总行数。
	TotalRows *int32 `json:"totalRows,omitempty"`

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	// 分页大小（路径参数）。
	PageSizePath int32 `json:"pageSizePath"`

	// 当前页数（路径参数）。
	CurPagePath int32 `json:"curPagePath"`

	Body *RdmParamVoQueryChildListDto `json:"body,omitempty"`
}

func (o ListGetChildListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGetChildListRequest struct{}"
	}

	return strings.Join([]string{"ListGetChildListRequest", string(data)}, " ")
}
