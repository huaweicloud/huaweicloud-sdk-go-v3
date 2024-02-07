package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalEipSegmentsResponse Response Object
type ListGlobalEipSegmentsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 全域弹性公网IP段对象
	GlobalEipSegments *[]ListGlobalEipSegments `json:"global_eip_segments,omitempty"`

	PageInfo *ListGlobalEipsResponseBodyPageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGlobalEipSegmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipSegmentsResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalEipSegmentsResponse", string(data)}, " ")
}
