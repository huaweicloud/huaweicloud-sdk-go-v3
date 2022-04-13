package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunSegmentRequest struct {
	Body *SegmentRequest `json:"body,omitempty"`
}

func (o RunSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSegmentRequest struct{}"
	}

	return strings.Join([]string{"RunSegmentRequest", string(data)}, " ")
}
