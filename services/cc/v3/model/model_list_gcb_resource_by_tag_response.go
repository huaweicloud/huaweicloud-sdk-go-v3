package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGcbResourceByTagResponse Response Object
type ListGcbResourceByTagResponse struct {

	// 资源列表。
	Resources *[]TmsResource `json:"resources,omitempty"`

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGcbResourceByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGcbResourceByTagResponse struct{}"
	}

	return strings.Join([]string{"ListGcbResourceByTagResponse", string(data)}, " ")
}
