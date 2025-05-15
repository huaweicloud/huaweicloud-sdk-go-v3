package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoIdDocClassificationAlarmConfidence struct {

	// 证件图像模糊告警分数，分数越高，证件图像模糊的可能性越高。
	BlurScore *int32 `json:"blur_score,omitempty"`

	// 证件图像反光告警分数，分数越高，证件图像反光的可能性越高。
	GlareScore *int32 `json:"glare_score,omitempty"`

	// 证件图像框内遮挡告警分数，分数越高，证件图像框内遮挡的可能性越高。
	BlockingWithinBorderScore *int32 `json:"blocking_within_border_score,omitempty"`

	// 证件图像过暗告警分数，分数越高，证件图像过暗的可能性越高。
	InsufficientLightingScore *int32 `json:"insufficient_lighting_score,omitempty"`

	// 证件图像复印件告警分数，分数越高，证件图像是复印件的可能性越高。
	CopyScore *int32 `json:"copy_score,omitempty"`

	// 证件图像边框完整性告警分数，分数越高，证件图像边框不完整的可能性越高。
	BorderIntegrityScore *int32 `json:"border_integrity_score,omitempty"`

	// 证件图像翻拍告警分数，分数越高，证件图像经过翻拍的可能性越高。
	ReproduceScore *int32 `json:"reproduce_score,omitempty"`
}

func (o AutoIdDocClassificationAlarmConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoIdDocClassificationAlarmConfidence struct{}"
	}

	return strings.Join([]string{"AutoIdDocClassificationAlarmConfidence", string(data)}, " ")
}
