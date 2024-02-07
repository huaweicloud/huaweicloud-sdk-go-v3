package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeipSegmentFilterTagsResponse Response Object
type ListGeipSegmentFilterTagsResponse struct {

	// 资源列表
	Resources *[]GeipResource `json:"resources,omitempty"`

	// 当前列表中资源数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGeipSegmentFilterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipSegmentFilterTagsResponse struct{}"
	}

	return strings.Join([]string{"ListGeipSegmentFilterTagsResponse", string(data)}, " ")
}
