package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageSuperResolutionResponse Response Object
type RunImageSuperResolutionResponse struct {
	Result         *ImageSuperResolutionResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o RunImageSuperResolutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageSuperResolutionResponse struct{}"
	}

	return strings.Join([]string{"RunImageSuperResolutionResponse", string(data)}, " ")
}
