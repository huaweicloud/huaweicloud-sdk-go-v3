package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdcardBackResult
type IdcardBackResult struct {

	// 发证机关。
	Issue *string `json:"issue,omitempty"`

	// 有效起始日期。
	ValidFrom *string `json:"valid_from,omitempty"`

	// 有效结束日期。
	ValidTo *string `json:"valid_to,omitempty"`

	// 身份证卡面图片信息的base64码结果。  > 说明： - 仅在输入参数return_adjusted_image为true时，返回该字段。
	AdjustedImage *string `json:"adjusted_image,omitempty"`

	VerificationResult *IdcardBackVerificationResult `json:"verification_result,omitempty"`

	// 文本框在原图位置。输出左上、右上、右下、左下四个点坐标。 仅return_text_location设置为true时才返回。
	TextLocation *interface{} `json:"text_location,omitempty"`

	// 身份证图像是否翻拍告警结果。 - true：表示身份证图片经过翻拍。 - false：表示身份证图片未经过翻拍。 仅在输入参数detect_reproduce为true时，返回该字段。
	DetectReproduceResult *bool `json:"detect_reproduce_result,omitempty"`

	// 身份证图像是否黑白复印件告警结果。 - true：表示身份证图片是复印件。 - false”表示身份证图片是原件。 仅在输入参数detect_copy为true时，返回该字段。
	DetectCopyResult *bool `json:"detect_copy_result,omitempty"`

	// 身份证图片是否PS告警结果。 - true：表示身份证经过PS。 - false：表示未经过PS。 仅在传入参数detect_tampering为true时，返回该字段。
	DetectTamperingResult *bool `json:"detect_tampering_result,omitempty"`

	// 身份证图片边框完整性告警结果。 - true：表示边框不完整 - false：表示边框完整。 仅在输入参数detect_border_integrity为true时，返回该字段。
	DetectBorderIntegrityResult *bool `json:"detect_border_integrity_result,omitempty"`

	// 身份证图像框内是否存在遮挡的告警结果。 - true：表示边框内部存在遮挡。 - false：表示边框内部不存在遮挡。 仅在输入参数detect_blocking_within_border为true时，返回该字段。
	DetectBlockingWithinBorderResult *bool `json:"detect_blocking_within_border_result,omitempty"`

	// 身份证模糊告警结果。 - true：表示身份证图片较模糊。 - false：表示身份证清晰。 仅在输入参数detect_blur为true时，返回该字段。
	DetectBlurResult *bool `json:"detect_blur_result,omitempty"`

	// 临时身份证告警结果。 - true：表示是临时身份证。 - false：表示非临时身份证。 仅在输入参数detect_interim为true时，返回该字段。
	DetectInterimResult *bool `json:"detect_interim_result,omitempty"`

	// 身份证反光告警结果。 - true：表示身份证图片存在反光。 - false：表示是身份证不存在反光。 仅在输入参数detect_glare为true时，返回该字段。
	DetectGlareResult *bool `json:"detect_glare_result,omitempty"`

	ScoreInfo *IdcardScoreInfoResult `json:"score_info,omitempty"`
}

func (o IdcardBackResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdcardBackResult struct{}"
	}

	return strings.Join([]string{"IdcardBackResult", string(data)}, " ")
}
