package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllVersionsRequest Request Object
type ListAllVersionsRequest struct {

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

	// 当前页
	CurPagePath int32 `json:"curPagePath"`

	// 每页大小
	PageSizePath int32 `json:"pageSizePath"`

	Body *RdmParamVoVersionModelVersionMasterDto `json:"body,omitempty"`
}

func (o ListAllVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListAllVersionsRequest", string(data)}, " ")
}
