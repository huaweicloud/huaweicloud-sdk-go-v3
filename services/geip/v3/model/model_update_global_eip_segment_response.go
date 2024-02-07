package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalEipSegmentResponse Response Object
type UpdateGlobalEipSegmentResponse struct {

	// 本次请求的job id
	JobId *string `json:"job_id,omitempty"`

	GlobalEipSegment *ShortGlobalEipSegment `json:"global_eip_segment,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGlobalEipSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipSegmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipSegmentResponse", string(data)}, " ")
}
