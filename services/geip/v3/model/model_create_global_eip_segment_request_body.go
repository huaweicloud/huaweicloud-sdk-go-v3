package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGlobalEipSegmentRequestBody struct {
	GlobalEipSegment *CreateGlobalEipSegmentRequestBodyGlobalEipSegment `json:"global_eip_segment"`
}

func (o CreateGlobalEipSegmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipSegmentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipSegmentRequestBody", string(data)}, " ")
}
