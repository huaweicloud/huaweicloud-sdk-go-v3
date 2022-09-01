package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunImageMainObjectDetectionRequest struct {
	Body *ImageMainObjectDetectionReq `json:"body,omitempty" xml:"body"`
}

func (o RunImageMainObjectDetectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageMainObjectDetectionRequest struct{}"
	}

	return strings.Join([]string{"RunImageMainObjectDetectionRequest", string(data)}, " ")
}
