package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalEipSegmentResponse Response Object
type ShowGlobalEipSegmentResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	GlobalEipSegment *ShowGlobalEipSegment `json:"global_eip_segment,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGlobalEipSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalEipSegmentResponse struct{}"
	}

	return strings.Join([]string{"ShowGlobalEipSegmentResponse", string(data)}, " ")
}
