package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateGeipSegmentInstanceResponse Response Object
type DisassociateGeipSegmentInstanceResponse struct {

	// 本次请求的job id
	JobId *string `json:"job_id,omitempty"`

	GlobalEipSegment *ShortGlobalEipSegment `json:"global_eip_segment,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateGeipSegmentInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateGeipSegmentInstanceResponse struct{}"
	}

	return strings.Join([]string{"DisassociateGeipSegmentInstanceResponse", string(data)}, " ")
}
