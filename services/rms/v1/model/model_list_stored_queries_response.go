package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStoredQueriesResponse struct {

	// 高级查询列表
	Value *[]StoredQuery `json:"value,omitempty" xml:"value"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStoredQueriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoredQueriesResponse struct{}"
	}

	return strings.Join([]string{"ListStoredQueriesResponse", string(data)}, " ")
}
