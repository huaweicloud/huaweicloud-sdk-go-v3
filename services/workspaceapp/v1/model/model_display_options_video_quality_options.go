package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisplayOptionsVideoQualityOptions 视频质量控制选项。
type DisplayOptionsVideoQualityOptions struct {

	// 视频平均质量。取值范围为[5-59]。默认：15。
	AverageVideoQuality *int32 `json:"average_video_quality,omitempty"`

	// 视频最低质量。取值范围为[5-69]。默认：25。
	LowestVideoQuality *int32 `json:"lowest_video_quality,omitempty"`

	// 视频最高质量。取值范围为[1-59]。默认：7。
	HighestVideoQuality *int32 `json:"highest_video_quality,omitempty"`
}

func (o DisplayOptionsVideoQualityOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayOptionsVideoQualityOptions struct{}"
	}

	return strings.Join([]string{"DisplayOptionsVideoQualityOptions", string(data)}, " ")
}
