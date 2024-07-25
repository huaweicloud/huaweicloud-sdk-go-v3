package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CambodianIdCardScoreInformationResult struct {

	// 告警分数，字段取值范围[0, 99]值大于50表示复印件，小于等于50表示原件，值越靠近99，表示复印件的可能性越大，值越靠近0，表示原件的可能性越大。  仅在传入参数return_idcard_type为true时，返回该字段。
	IdcardTypeScore *int32 `json:"idcard_type_score,omitempty"`

	// 告警分数，字段取值范围[0, 99]值大于50表示边框不完整，小于50表示边框完整，值越靠近99，表示边框不完整的可能性越大，值越靠近0，表示边框完整的可能性越大。 仅在传入参数detect_border_integrity为true时，返回该字段。
	BorderIntegrityScore *int32 `json:"border_integrity_score,omitempty"`

	// 告警分数，字段取值范围[0, 99]值大于50表示框内有遮挡，小于50表示框内无遮挡，值越靠近99，表示框内有遮挡的可能性越大，值越靠近0，表示框内无遮挡的可能性越大。 仅在传入参数detect_blocking_within_border为true时，返回该字段。
	BlockingWithinBorderScore *int32 `json:"blocking_within_border_score,omitempty"`

	// 告警分数，字段取值范围[0, 99]值大于50表示身份证模糊，小于50表示身份证清晰，值越靠近99，表示身份证模糊的可能性越大，值越靠近0，表示身份证清晰的可能性越大。 仅在传入参数detect_blur为true时，返回该字段。
	BlurScore *int32 `json:"blur_score,omitempty"`

	// 告警分数，字段取值范围[0, 99]值大于50表示身份证反光，小于50表示身份证不反光，值越靠近99，表示身份证反光的可能性越大，值越靠近0，表示身份证不反光的可能性越大。 仅在传入参数detect_glare为true时，返回该字段。
	GlareScore *int32 `json:"glare_score,omitempty"`

	// 告警分数，字段取值范围[0, 99]值大于50表示身份证人像被其他非身份证人像篡改过，小于50表示身份证人像未被篡改，值越靠近99，表示身份证人像被篡改的可能性越大，值越靠近0，表示身份证未人像被篡改的可能性越大。 仅在传入参数detect_tampering为true时，返回该字段。
	TamperingScore *int32 `json:"tampering_score,omitempty"`

	// 告警分数，字段取值范围[0, 99]值大于50表示身份证经过翻拍，小于50表示身份证未经过翻拍，值越靠近99，表示身份证图像被翻拍过的可能性越大，值越靠近0，表示身份证图像未被翻拍的可能性越大。 仅在传入参数detect_reproduce为true时，返回该字段。
	ReproduceScore *int32 `json:"reproduce_score,omitempty"`
}

func (o CambodianIdCardScoreInformationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CambodianIdCardScoreInformationResult struct{}"
	}

	return strings.Join([]string{"CambodianIdCardScoreInformationResult", string(data)}, " ")
}
