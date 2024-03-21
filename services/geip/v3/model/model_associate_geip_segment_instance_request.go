package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateGeipSegmentInstanceRequest Request Object
type AssociateGeipSegmentInstanceRequest struct {
	GlobalEipSegmentId string `json:"global_eip_segment_id"`

	Body *AssociateInstanceGlobalEipSegmentRequestBody `json:"body,omitempty"`
}

func (o AssociateGeipSegmentInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGeipSegmentInstanceRequest struct{}"
	}

	return strings.Join([]string{"AssociateGeipSegmentInstanceRequest", string(data)}, " ")
}
