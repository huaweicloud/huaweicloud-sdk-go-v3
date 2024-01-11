package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdcardScoreInfoResult
type IdcardScoreInfoResult struct {

	// 身份证PS告警分数，分数越高，PS的可能性越高。 仅在传入参数detect_tampering为true时，返回该字段。
	TamperingScore *int32 `json:"tampering_score,omitempty"`

	// 临时身份证告警分数，分数越高，临时身份证的可能性越高。 仅在传入参数detect_interim为true时，返回该字段。
	InterimScore *int32 `json:"interim_score,omitempty"`

	// 身份证翻拍告警分数，分数越高，身份证图像经过翻拍的可能性越高。 仅在传入参数detect_reproduce为true时，返回该字段。
	ReproduceScore *int32 `json:"reproduce_score,omitempty"`

	// 身份证复印告警分数, 分数越高，身份证图像是复印件的可能性越高。 仅在传入参数detect_copy为true时，返回该字段。
	CopyScore *int32 `json:"copy_score,omitempty"`

	// 身份证边缘完整性告警的分数，分数越高，身份证图像边缘不完整的可能性越高。 仅在传入参数detect_border_integrity为true时，返回该字段。
	BorderIntegrityScore *int32 `json:"border_integrity_score,omitempty"`

	// 身份证模糊告警分数，分数越高，身份证图像模糊的可能性越高。 仅在传入参数detect_blur为true时，返回该字段。
	BlurScore *int32 `json:"blur_score,omitempty"`

	// 身份证反光告警分数，分数越高，身份证图像存在反光的可能性越高。 仅在传入参数detect_glare为true时，返回该字段。
	GlareScore *int32 `json:"glare_score,omitempty"`

	// 身份证图像框内是否有遮挡的告警分数，分数越高，身份证图像框内存在异物遮挡的可能性越高。 仅在传入参数detect_blocking_within_border为true时，返回该字段。
	BlockingWithinBorderScore *int32 `json:"blocking_within_border_score,omitempty"`
}

func (o IdcardScoreInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdcardScoreInfoResult struct{}"
	}

	return strings.Join([]string{"IdcardScoreInfoResult", string(data)}, " ")
}
