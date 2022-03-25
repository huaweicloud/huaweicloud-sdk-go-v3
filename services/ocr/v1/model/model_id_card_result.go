package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type IdCardResult struct {
	// 姓名。

	Name *string `json:"name,omitempty"`
	// 性别。

	Sex *string `json:"sex,omitempty"`
	// 出生日期。

	Birth *string `json:"birth,omitempty"`
	// 民族。

	Ethnicity *string `json:"ethnicity,omitempty"`
	// 地址。

	Address *string `json:"address,omitempty"`
	// 身份证号。

	Number *string `json:"number,omitempty"`
	// 发证机关。

	Issue *string `json:"issue,omitempty"`
	// 有效起始日期。

	ValidFrom *string `json:"valid_from,omitempty"`
	// 有效结束日期。   > 说明：  - 身份证识别支持中华人民共和国居民身份证识别。

	ValidTo *string `json:"valid_to,omitempty"`

	VerificationResult *IdcardVerificationResult `json:"verification_result,omitempty"`
	// 文本框在原图位置。输出左上、右上、右下、左下四个点坐标。当“return_text_location”设置为“true”时才返回。

	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o IdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdCardResult struct{}"
	}

	return strings.Join([]string{"IdCardResult", string(data)}, " ")
}
