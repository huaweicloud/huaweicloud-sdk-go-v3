package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunImageWisedesignInpaintingResponse struct {
	Result         *ImageWisedesignInpaintingResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o RunImageWisedesignInpaintingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageWisedesignInpaintingResponse struct{}"
	}

	return strings.Join([]string{"RunImageWisedesignInpaintingResponse", string(data)}, " ")
}
