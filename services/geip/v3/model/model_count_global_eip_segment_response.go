package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountGlobalEipSegmentResponse Response Object
type CountGlobalEipSegmentResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	GlobalEipSegment *CountGeipSegments `json:"global_eip_segment,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountGlobalEipSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGlobalEipSegmentResponse struct{}"
	}

	return strings.Join([]string{"CountGlobalEipSegmentResponse", string(data)}, " ")
}
