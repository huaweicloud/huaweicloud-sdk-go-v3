package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeipSegmentDomainTagsResponse Response Object
type ListGeipSegmentDomainTagsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 当前列表中资源数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	// tag列表信息
	Tags *[]GetAccountTagsV2 `json:"tags,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGeipSegmentDomainTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipSegmentDomainTagsResponse struct{}"
	}

	return strings.Join([]string{"ListGeipSegmentDomainTagsResponse", string(data)}, " ")
}
