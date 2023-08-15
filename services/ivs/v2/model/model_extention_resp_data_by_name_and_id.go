package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtentionRespDataByNameAndId struct {

	// 审核校验结果： \"valid\"表示身份审核通过； \"invalid\"表示身份审核不通过； \"nonexistent\"表示数据源没有该身份证号码，这种情况一般是被验证人正在办理户籍迁移，或者被验证人是军人或政要。
	VerificationResult *string `json:"verification_result,omitempty"`

	// 审核校验信息，具体参考[校验信息说明](https://support.huaweicloud.com/api-ivs/ivs_02_0017.html)
	VerificationMessage *string `json:"verification_message,omitempty"`

	// 审核校验代码，具体参考[校验信息说明](https://support.huaweicloud.com/api-ivs/ivs_02_0017.html)
	VerificationCode *int32 `json:"verification_code,omitempty"`
}

func (o ExtentionRespDataByNameAndId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtentionRespDataByNameAndId struct{}"
	}

	return strings.Join([]string{"ExtentionRespDataByNameAndId", string(data)}, " ")
}
