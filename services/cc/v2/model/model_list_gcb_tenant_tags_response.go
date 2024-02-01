package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGcbTenantTagsResponse Response Object
type ListGcbTenantTagsResponse struct {

	// 标签列表。
	Tags *[]TmsTagValues `json:"tags,omitempty"`

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGcbTenantTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGcbTenantTagsResponse struct{}"
	}

	return strings.Join([]string{"ListGcbTenantTagsResponse", string(data)}, " ")
}
