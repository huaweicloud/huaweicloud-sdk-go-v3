package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalEipSegmentRequest Request Object
type CreateGlobalEipSegmentRequest struct {
	Body *CreateGlobalEipSegmentRequestBody `json:"body,omitempty"`
}

func (o CreateGlobalEipSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipSegmentRequest struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipSegmentRequest", string(data)}, " ")
}
