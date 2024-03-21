package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryDataRequest Request Object
type ListHistoryDataRequest struct {

	// 当前页。
	CurPage *int32 `json:"curPage,omitempty"`

	// 每页大小。
	PageSize *int32 `json:"pageSize,omitempty"`

	// 总行数。
	TotalRows *int32 `json:"totalRows,omitempty"`

	// 总页数。
	TotalPages *int32 `json:"totalPages,omitempty"`

	// 每页显示条目数量，limit和offset均传正确的数值时才起作用，且优先级高于pageSize和curPage。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，limit和offset均传正确的数值时才起作用，且优先级高于pageSize和curPage。
	Offset *int32 `json:"offset,omitempty"`

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	// 分页大小（路径参数）。
	PageSizePath int32 `json:"pageSizePath"`

	// 当前页数（路径参数）。
	CurPagePath int32 `json:"curPagePath"`

	Body *RdmParamVoMongPageRequest `json:"body,omitempty"`
}

func (o ListHistoryDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryDataRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryDataRequest", string(data)}, " ")
}
