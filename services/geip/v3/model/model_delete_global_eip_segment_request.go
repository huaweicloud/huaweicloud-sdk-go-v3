package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalEipSegmentRequest Request Object
type DeleteGlobalEipSegmentRequest struct {
	GlobalEipSegmentId string `json:"global_eip_segment_id"`
}

func (o DeleteGlobalEipSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalEipSegmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteGlobalEipSegmentRequest", string(data)}, " ")
}
