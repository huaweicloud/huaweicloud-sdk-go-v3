package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeystoreSearchRequest Request Object
type ListKeystoreSearchRequest struct {

	// 每页显示的条目数量，page_size小于等于100
	PageSize *int32 `json:"page_size,omitempty"`

	// 分页页码，表示从此页开始查询， page大于等于1
	Page *int32 `json:"page,omitempty"`

	// 搜索的文件名
	Search *string `json:"search,omitempty"`
}

func (o ListKeystoreSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeystoreSearchRequest struct{}"
	}

	return strings.Join([]string{"ListKeystoreSearchRequest", string(data)}, " ")
}
