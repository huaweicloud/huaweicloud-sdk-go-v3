package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 校验AIM能力消息体。
type CheckMobileCapabilityRequestBody struct {

	// 待查询的手机号码，一次最多请求100个。  >不加密时，参数可传入纯手机号或国家码加手机号，国家码不带“+”号，例如国内手机号“131****0000”，参数可传入“131****0000”、“86131****0000”、“0086131****0000”。使用SHA1加密，mobiles传入号码国家码与手机号码的SHA1算法后的摘要，国家码格式为纯数字，不带“+”，去掉前面的0，例如国内手机号“131****0000”，使用“86131****0000”进行SHA1加密。
	Mobiles []string `json:"mobiles"`

	// 智能信息模板ID，由9位数字组成。 > - 填写时，根据该模板所支持的厂商返回手机终端展示智能信息的能力状态 > - 不填则返回手机终端在所有厂商展示智能信息的能力状态
	TplId *string `json:"tpl_id,omitempty"`

	// 加密类型。  - NONE：不加密 - SHA1：使用SHA1加密算法加密  > 默认为NONE。
	EncryptionAlg *string `json:"encryption_alg,omitempty"`
}

func (o CheckMobileCapabilityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMobileCapabilityRequestBody struct{}"
	}

	return strings.Join([]string{"CheckMobileCapabilityRequestBody", string(data)}, " ")
}
