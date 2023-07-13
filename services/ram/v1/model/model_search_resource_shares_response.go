package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceSharesResponse Response Object
type SearchResourceSharesResponse struct {

	// 资源共享实例的详细信息列表。
	ResourceShares *[]ResourceShare `json:"resource_shares,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchResourceSharesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceSharesResponse struct{}"
	}

	return strings.Join([]string{"SearchResourceSharesResponse", string(data)}, " ")
}
