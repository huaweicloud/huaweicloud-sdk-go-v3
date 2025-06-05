package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecordsResult 项目列表
type ListRecordsResult struct {
	Pagination *ListRecordsResultPagination `json:"pagination,omitempty"`

	// 工程构建记录列表
	Data *[]BuildRecord `json:"data,omitempty"`
}

func (o ListRecordsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordsResult struct{}"
	}

	return strings.Join([]string{"ListRecordsResult", string(data)}, " ")
}
