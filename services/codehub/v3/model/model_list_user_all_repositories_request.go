package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListUserAllRepositoriesRequest struct {

	// 分页索引，从1开始计数
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index"`

	// 每页条目数
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 搜索关键字
	Search *string `json:"search,omitempty" xml:"search"`
}

func (o ListUserAllRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserAllRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListUserAllRepositoriesRequest", string(data)}, " ")
}
