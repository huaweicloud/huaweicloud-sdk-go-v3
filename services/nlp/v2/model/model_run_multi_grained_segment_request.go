package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunMultiGrainedSegmentRequest struct {
	Body *PostMultiGrainedSegmentReq `json:"body,omitempty"`
}

func (o RunMultiGrainedSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunMultiGrainedSegmentRequest struct{}"
	}

	return strings.Join([]string{"RunMultiGrainedSegmentRequest", string(data)}, " ")
}
