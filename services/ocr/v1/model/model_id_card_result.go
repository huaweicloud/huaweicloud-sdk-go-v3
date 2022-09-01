package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type IdCardResult struct {

	// 姓名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 性别。
	Sex *string `json:"sex,omitempty" xml:"sex"`

	// 出生日期。
	Birth *string `json:"birth,omitempty" xml:"birth"`

	// 民族。
	Ethnicity *string `json:"ethnicity,omitempty" xml:"ethnicity"`

	// 地址。
	Address *string `json:"address,omitempty" xml:"address"`

	// 身份证号。
	Number *string `json:"number,omitempty" xml:"number"`

	// 发证机关。
	Issue *string `json:"issue,omitempty" xml:"issue"`

	// 有效起始日期。
	ValidFrom *string `json:"valid_from,omitempty" xml:"valid_from"`

	// 有效结束日期。   > 说明：  - 身份证识别支持中华人民共和国居民身份证识别。
	ValidTo *string `json:"valid_to,omitempty" xml:"valid_to"`

	VerificationResult *IdcardVerificationResult `json:"verification_result,omitempty" xml:"verification_result"`

	// 文本框在原图位置。输出左上、右上、右下、左下四个点坐标。当“return_text_location”设置为“true”时才返回。
	TextLocation *interface{} `json:"text_location,omitempty" xml:"text_location"`

	// 判断身份证图像是否经过翻拍，“true”表示是翻拍，“false”表示未经过翻拍。仅在输入参数detect_reproduce为true时，返回该字段。
	DetectReproduceResult *bool `json:"detect_reproduce_result,omitempty" xml:"detect_reproduce_result"`
}

func (o IdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdCardResult struct{}"
	}

	return strings.Join([]string{"IdCardResult", string(data)}, " ")
}
