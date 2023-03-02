package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunImageWisedesignInpaintingRequest struct {
	Body *ImageWisedesignInpaintingReq `json:"body,omitempty"`
}

func (o RunImageWisedesignInpaintingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageWisedesignInpaintingRequest struct{}"
	}

	return strings.Join([]string{"RunImageWisedesignInpaintingRequest", string(data)}, " ")
}
