package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageSuperResolutionRequest Request Object
type RunImageSuperResolutionRequest struct {
	Body *ImageSuperResolutionReq `json:"body,omitempty"`
}

func (o RunImageSuperResolutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageSuperResolutionRequest struct{}"
	}

	return strings.Join([]string{"RunImageSuperResolutionRequest", string(data)}, " ")
}
