package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecordsResultPagination 分页信息
type ListRecordsResultPagination struct {

	// 分页页数
	Page *int32 `json:"page,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`
}

func (o ListRecordsResultPagination) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordsResultPagination struct{}"
	}

	return strings.Join([]string{"ListRecordsResultPagination", string(data)}, " ")
}
