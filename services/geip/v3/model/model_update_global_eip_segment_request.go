package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalEipSegmentRequest Request Object
type UpdateGlobalEipSegmentRequest struct {
	GlobalEipSegmentId string `json:"global_eip_segment_id"`

	Body *UpdateGlobalEipSegmentRequestBody `json:"body,omitempty"`
}

func (o UpdateGlobalEipSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipSegmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipSegmentRequest", string(data)}, " ")
}
