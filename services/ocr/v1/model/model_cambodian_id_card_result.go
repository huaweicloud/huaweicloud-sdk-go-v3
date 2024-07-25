package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CambodianIdCardResult struct {

	// 身份证号码。
	IdNumber *string `json:"id_number,omitempty"`

	// 高棉语版姓名。
	NameKh *string `json:"name_kh,omitempty"`

	// 英文姓名。
	NameEn *string `json:"name_en,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 身高。
	Height *string `json:"height,omitempty"`

	// 出生地。
	BirthPlace *string `json:"birth_place,omitempty"`

	// 地址，以空格分隔。
	Address *string `json:"address,omitempty"`

	// 签发起始日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 签发结束日期。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 图片中的个人特征。
	Description *string `json:"description,omitempty"`

	// 机器码第一行。
	MachineCode1 *string `json:"machine_code1,omitempty"`

	// 机器码第二行。
	MachineCode2 *string `json:"machine_code2,omitempty"`

	// 机器码第三行。
	MachineCode3 *string `json:"machine_code3,omitempty"`

	// 头像的base64编码。 当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty"`

	// 身份证的类型。当输入参数\"idcard_type \"为\"true\"时，才返回该参数。取值如下所示： - normal：身份证原件 - copy：复印的身份证
	IdcardType *string `json:"idcard_type,omitempty"`

	// 身份证原图的base64编码。 当输入参数\"return_adjusted_image\"为\"true\"时，才返回该参数。
	AdjustedImage *string `json:"adjusted_image,omitempty"`

	// 身份证图片边框完整性告警结果，\"true\"表示边框不完整，\"false\"表示边框完整。仅在输入参数detect_border_integrity为true时，返回该字段。
	DetectBorderIntegrityResult *bool `json:"detect_border_integrity_result,omitempty"`

	// 身份证图像框内是否存在遮挡的告警结果，\"true\"表示边框内部存在遮挡，\"false\"表示边框内部完整。仅在输入参数detect_blocking_within_border为true时，返回该字段。
	DetectBlockingWithinBorderResult *bool `json:"detect_blocking_within_border_result,omitempty"`

	// 身份证模糊告警结果，\"true\"表示图片模糊，\"false\"表示身份证清晰。仅在输入参数detect_blur为true时，返回该字段。
	DetectBlurResult *bool `json:"detect_blur_result,omitempty"`

	// 身份证反光告警结果，\"true\"表示身份证反光，\"false\"表示是身份证无反光。仅在输入参数detect_glare为true时，返回该字段。
	DetectGlareResult *bool `json:"detect_glare_result,omitempty"`

	// 身份证人像被篡改的告警结果，\"true\"表示身份证人像被篡改，\"false\"表示是身份证人像未被篡改。仅在输入参数detect_tampering为true时，返回该字段。
	DetectTamperingResult *bool `json:"detect_tampering_result,omitempty"`

	// 身份证是否经过翻拍的告警结果，“true”表示身份证经过翻拍，“false”表示身份证未经过翻拍。仅在输入参数detect_reproduce为true时，返回该字段。
	DetectReproduceResult *bool `json:"detect_reproduce_result,omitempty"`

	ScoreInfo *CambodianIdCardScoreInformationResult `json:"score_info,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o CambodianIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CambodianIdCardResult struct{}"
	}

	return strings.Join([]string{"CambodianIdCardResult", string(data)}, " ")
}
