package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdcardVerificationResult 校验信息。仅在输入参数return_verification为true时，返回该字段，该字段包含valid_number、valid_birth、valid_sex和valid_date的校验信息。
type IdcardVerificationResult struct {

	// 身份证号规则校验是否通过。“true”表示身份证号规则校验通过，“false”表示身份证号规则校验不通过。当身份证图片是国徽面时，默认是false。仅在输入参数return_verification为true时，返回该字段。
	ValidNumber *bool `json:"valid_number,omitempty"`

	// 出生日期与身份证号所表示的出生日期是否一致。“true”表示一致，“false”表示不一致。当身份证图片是国徽面，或者身份证号规则校验不通过时，默认是false。仅在输入参数return_verification为true时，返回该字段。
	ValidBirth *bool `json:"valid_birth,omitempty"`

	// 性别与身份证号所表示的性别信息是否一致。“true”表示一致，“false”表示不一致。当身份证图片是国徽面，或者身份证号规则校验不通过时，默认是false。仅在输入参数return_verification为true时，返回该字段。
	ValidSex *bool `json:"valid_sex,omitempty"`

	// 当前日期是否在有效期内。“true”表示当前日期在有效期内，“false”表示当前日期不在有效期内。当身份证图片是人像面时，默认是false。仅在输入参数return_verification为true时，返回该字段。
	ValidDate *bool `json:"valid_date,omitempty"`
}

func (o IdcardVerificationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdcardVerificationResult struct{}"
	}

	return strings.Join([]string{"IdcardVerificationResult", string(data)}, " ")
}
