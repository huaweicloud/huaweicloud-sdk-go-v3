package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSegmentRequest Request Object
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
