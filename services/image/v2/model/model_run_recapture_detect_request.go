package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunRecaptureDetectRequest Request Object
type RunRecaptureDetectRequest struct {
	Body *RecaptureDetectReq `json:"body,omitempty"`
}

func (o RunRecaptureDetectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunRecaptureDetectRequest struct{}"
	}

	return strings.Join([]string{"RunRecaptureDetectRequest", string(data)}, " ")
}
