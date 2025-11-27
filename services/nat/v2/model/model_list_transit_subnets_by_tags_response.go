package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitSubnetsByTagsResponse Response Object
type ListTransitSubnetsByTagsResponse struct {

	// 资源列表。
	Resources *[]Resource `json:"resources,omitempty"`

	// 请求id。
	RequestId *string `json:"request_id,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTransitSubnetsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitSubnetsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTransitSubnetsByTagsResponse", string(data)}, " ")
}
