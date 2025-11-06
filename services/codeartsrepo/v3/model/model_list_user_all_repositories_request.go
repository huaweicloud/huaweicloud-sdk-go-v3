package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserAllRepositoriesRequest Request Object
type ListUserAllRepositoriesRequest struct {

	// 分页索引，从1开始计数
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页条目数
	PageSize *int32 `json:"page_size,omitempty"`

	// 搜索关键字
	Search *string `json:"search,omitempty"`
}

func (o ListUserAllRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserAllRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListUserAllRepositoriesRequest", string(data)}, " ")
}
