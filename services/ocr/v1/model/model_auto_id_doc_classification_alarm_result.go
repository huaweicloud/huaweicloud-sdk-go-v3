package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoIdDocClassificationAlarmResult struct {

	// 证件图像模糊告警结果。 - true：表示证件图片较模糊。 - false：表示证件清晰。
	DetectBlurResult *bool `json:"detect_blur_result,omitempty"`

	// 证件图像反光告警结果。 - true：表示证件图片存在反光。 - false：表示证件图片不存在反光。
	DetectGlareResult *bool `json:"detect_glare_result,omitempty"`

	// 证件图像框内遮挡告警结果。 - true：表示证件图片存在框内遮挡。 - false：表示证件图片不存在框内遮挡。
	DetectBlockingWithinBorderResult *bool `json:"detect_blocking_within_border_result,omitempty"`

	// 证件图像过暗告警结果。 - true：表示证件图片过暗。 - false：表示证件图片光线正常。
	DetectInsufficientLightingResult *bool `json:"detect_insufficient_lighting_result,omitempty"`

	// 证件图像是否黑白复印件告警结果。 - true：表示证件是复印件。 - false：表示证件是原件。
	DetectCopyResult *bool `json:"detect_copy_result,omitempty"`

	// 证件图像边框完整性告警结果。 - true：表示边框不完整。 - false：表示边框完整。
	DetectBorderIntegrityResult *bool `json:"detect_border_integrity_result,omitempty"`

	// 证件图像是否翻拍告警结果。 - true：表示证件图片经过翻拍。 - false：表示证件图片未经过翻拍。
	DetectReproduceResult *bool `json:"detect_reproduce_result,omitempty"`
}

func (o AutoIdDocClassificationAlarmResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoIdDocClassificationAlarmResult struct{}"
	}

	return strings.Join([]string{"AutoIdDocClassificationAlarmResult", string(data)}, " ")
}
