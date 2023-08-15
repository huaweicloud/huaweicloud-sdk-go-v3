package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoredQueriesResponse Response Object
type ListStoredQueriesResponse struct {

	// 高级查询列表
	Value *[]StoredQuery `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStoredQueriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoredQueriesResponse struct{}"
	}

	return strings.Join([]string{"ListStoredQueriesResponse", string(data)}, " ")
}
