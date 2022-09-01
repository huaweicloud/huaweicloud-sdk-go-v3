package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceQuality struct {

	// 人脸质量总分，取值范围[0-1]，分值越大质量越高。
	TotalScore float64 `json:"total_score" xml:"total_score"`

	// 模糊度，取值范围[0-1]，分值越大模糊问题越严重。
	Blur float64 `json:"blur" xml:"blur"`

	// 姿态，取值范围[0-1]，分值越大姿态问题越严重。
	Pose float64 `json:"pose" xml:"pose"`

	// 遮挡，取值范围[0-1]，分值越大遮挡问题越严重。
	Occlusion float64 `json:"occlusion" xml:"occlusion"`

	// 光照，取值范围[0-1]，分值越大光照问题越严重。
	Illumination float64 `json:"illumination" xml:"illumination"`
}

func (o FaceQuality) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceQuality struct{}"
	}

	return strings.Join([]string{"FaceQuality", string(data)}, " ")
}
