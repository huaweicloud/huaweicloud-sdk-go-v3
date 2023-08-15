package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagResourceInstancesResponse Response Object
type ListTagResourceInstancesResponse struct {

	// 资源列表。
	Resources *[]Resource `json:"resources,omitempty"`

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTagResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListTagResourceInstancesResponse", string(data)}, " ")
}
