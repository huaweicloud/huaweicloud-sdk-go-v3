package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunImageWisedesignCropRequest struct {
	Body *ImageWisedesignCropReq `json:"body,omitempty"`
}

func (o RunImageWisedesignCropRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageWisedesignCropRequest struct{}"
	}

	return strings.Join([]string{"RunImageWisedesignCropRequest", string(data)}, " ")
}
