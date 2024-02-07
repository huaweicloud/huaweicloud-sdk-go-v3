package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGlobalEipSegmentRequestBody struct {
	GlobalEipSegment *UpdateGlobalEipSegmentRequestBodyGlobalEipSegment `json:"global_eip_segment"`
}

func (o UpdateGlobalEipSegmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipSegmentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipSegmentRequestBody", string(data)}, " ")
}
