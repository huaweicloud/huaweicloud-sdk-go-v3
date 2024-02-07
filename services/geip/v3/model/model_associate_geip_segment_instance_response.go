package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateGeipSegmentInstanceResponse Response Object
type AssociateGeipSegmentInstanceResponse struct {

	// 本次请求的job id
	JobId *string `json:"job_id,omitempty"`

	GlobalEipSegment *ShortGlobalEipSegment `json:"global_eip_segment,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateGeipSegmentInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGeipSegmentInstanceResponse struct{}"
	}

	return strings.Join([]string{"AssociateGeipSegmentInstanceResponse", string(data)}, " ")
}
