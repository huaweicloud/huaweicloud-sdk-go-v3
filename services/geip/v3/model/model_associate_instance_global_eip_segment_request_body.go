package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateInstanceGlobalEipSegmentRequestBody struct {
	GlobalEipSegment *AssociateInstanceGlobalEipSegmentRequestBodyGlobalEipSegment `json:"global_eip_segment"`
}

func (o AssociateInstanceGlobalEipSegmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceGlobalEipSegmentRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateInstanceGlobalEipSegmentRequestBody", string(data)}, " ")
}
