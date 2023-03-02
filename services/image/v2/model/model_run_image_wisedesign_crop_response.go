package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunImageWisedesignCropResponse struct {
	Result         *ImageWisedesignCropResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o RunImageWisedesignCropResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageWisedesignCropResponse struct{}"
	}

	return strings.Join([]string{"RunImageWisedesignCropResponse", string(data)}, " ")
}
