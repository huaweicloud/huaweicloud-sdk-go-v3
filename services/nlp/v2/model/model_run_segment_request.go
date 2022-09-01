package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunSegmentRequest struct {
	Body *SegmentRequest `json:"body,omitempty" xml:"body"`
}

func (o RunSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSegmentRequest struct{}"
	}

	return strings.Join([]string{"RunSegmentRequest", string(data)}, " ")
}
