package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecyclingJobRequest Request Object
type ListRecyclingJobRequest struct {

	// 每页显示的条目数量，page_size小于等于100
	PageSize *int32 `json:"page_size,omitempty"`

	// 分页页码， 表示从此页开始查询， page_no大于等于1
	PageNo *int32 `json:"page_no,omitempty"`

	// 搜索的任务名称
	Search *string `json:"search,omitempty"`
}

func (o ListRecyclingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecyclingJobRequest struct{}"
	}

	return strings.Join([]string{"ListRecyclingJobRequest", string(data)}, " ")
}
