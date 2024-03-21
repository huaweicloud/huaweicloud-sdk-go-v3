package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateGeipSegmentInstanceRequest Request Object
type DisassociateGeipSegmentInstanceRequest struct {
	GlobalEipSegmentId string `json:"global_eip_segment_id"`
}

func (o DisassociateGeipSegmentInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateGeipSegmentInstanceRequest struct{}"
	}

	return strings.Join([]string{"DisassociateGeipSegmentInstanceRequest", string(data)}, " ")
}
