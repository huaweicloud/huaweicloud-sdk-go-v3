package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelatedProjectInfoRequest Request Object
type ListRelatedProjectInfoRequest struct {

	// 每页显示的条目数量，page_size小于等于100
	PageSize *int32 `json:"page_size,omitempty"`

	// 分页页码，表示从此页开始查询， page大于等于1
	PageNo *int32 `json:"page_no,omitempty"`

	// 搜索内容
	Search *string `json:"search,omitempty"`
}

func (o ListRelatedProjectInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedProjectInfoRequest struct{}"
	}

	return strings.Join([]string{"ListRelatedProjectInfoRequest", string(data)}, " ")
}
