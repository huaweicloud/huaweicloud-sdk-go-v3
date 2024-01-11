package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdcardFrontVerificationResult
type IdcardFrontVerificationResult struct {

	// 身份证号规则校验是否通过。 - true：表示身份证号规则校验通过。 - false：表示身份证号规则校验不通过。 当识别结果为单页，身份证图片是国徽面时，默认是false。输入参数side为double_side时，该字典仅在front字段中存在。 仅在输入参数return_verification为true时，返回该字段。
	ValidNumber *bool `json:"valid_number,omitempty"`

	// 出生日期与身份证号所表示的出生日期是否一致。 - true：出生日期与身份证号所表示的出生日期一致。 - false：出生日期与身份证号所表示的出生日期不一致。 当识别结果为单页，身份证图片是国徽面，或者身份证号规则校验不通过时，默认是false。输入参数side为double_side时，该字段仅在front字典中存在。 仅在输入参数return_verification为true时，返回该字段。
	ValidBirth *bool `json:"valid_birth,omitempty"`

	// 性别与身份证号所表示的性别信息是否一致。 -true：性别与身份证号所表示的性别信息一致 -false：性别与身份证号所表示的性别信息不一致。 当识别结果为单页，身份证图片是国徽面，或者身份证号规则校验不通过时，默认是false。输入参数side为double_side时，该字段仅在front字典中存在。 仅在输入参数return_verification为true时，返回该字段。
	ValidSex *bool `json:"valid_sex,omitempty"`
}

func (o IdcardFrontVerificationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdcardFrontVerificationResult struct{}"
	}

	return strings.Join([]string{"IdcardFrontVerificationResult", string(data)}, " ")
}
