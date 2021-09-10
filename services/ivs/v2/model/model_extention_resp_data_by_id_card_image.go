package model

import (
	"encoding/json"

	"strings"
)

type ExtentionRespDataByIdCardImage struct {
	// 审核校验结果： \"valid\"表示身份审核通过； \"invalid\"表示身份审核不通过； \"nonexistent\"表示数据源没有该身份证号码，这种情况一般是被验证人正在办理户籍迁移，或者被验证人是军人或政要。

	VerificationResult *string `json:"verification_result,omitempty"`
	// 审核校验信息，具体参[考校验信息说明](https://support.huaweicloud.com/api-ivs/ivs_02_0017.html)

	VerificationMessage *string `json:"verification_message,omitempty"`
	// 审核校验代码，具体参[考校验信息说明](https://support.huaweicloud.com/api-ivs/ivs_02_0017.html)

	VerificationCode *int32 `json:"verification_code,omitempty"`

	IdcardResult *IdcardResult `json:"idcard_result,omitempty"`
}

func (o ExtentionRespDataByIdCardImage) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExtentionRespDataByIdCardImage struct{}"
	}

	return strings.Join([]string{"ExtentionRespDataByIdCardImage", string(data)}, " ")
}
