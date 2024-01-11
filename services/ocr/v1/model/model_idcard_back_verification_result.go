package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdcardBackVerificationResult
type IdcardBackVerificationResult struct {

	// 当前日期是否在有效期内。 - true：表示当前日期在有效期内。 - false：表示当前日期不在有效期内。 当识别结果为单页，身份证图片是人像面时，默认是false。输入参数side为double_side时，该字段仅在back字典中存在。 仅在输入参数return_verification为true时，返回该字段。
	ValidDate *bool `json:"valid_date,omitempty"`

	// 身份证有效日期是否合法。 - true：表示身份证的有效日期合法 - false：表示身份证有效日期非法 当识别结果为单页，身份证图片是人像面时，默认是false。输入参数side为double_side时，该字段仅在back字典中存在。 仅在输入参数return_verification为true时，返回该字段。
	ValidValidityPeriod *bool `json:"valid_validity_period,omitempty"`
}

func (o IdcardBackVerificationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdcardBackVerificationResult struct{}"
	}

	return strings.Join([]string{"IdcardBackVerificationResult", string(data)}, " ")
}
