package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualityEnhanceVideo struct {
	VideoDenoise *VideoDenoise `json:"video_denoise,omitempty" xml:"video_denoise"`

	VideoSharp *VideoSharp `json:"video_sharp,omitempty" xml:"video_sharp"`

	VideoContrast *VideoContrast `json:"video_contrast,omitempty" xml:"video_contrast"`

	VideoSuperresolution *VideoSuperresolution `json:"video_superresolution,omitempty" xml:"video_superresolution"`

	VideoDeblock *VideoDeblock `json:"video_deblock,omitempty" xml:"video_deblock"`

	VideoSaturation *VideoSaturation `json:"video_saturation,omitempty" xml:"video_saturation"`
}

func (o QualityEnhanceVideo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityEnhanceVideo struct{}"
	}

	return strings.Join([]string{"QualityEnhanceVideo", string(data)}, " ")
}
