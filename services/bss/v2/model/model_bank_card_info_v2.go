package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BankCardInfoV2 struct {

	// 银行卡账号。 当identifyType为4时，不能为空。 银行账号输入规则：^[0-9a-zA-Z]，可以包含特殊横杠（-）字符。
	BankAccount string `json:"bank_account"`

	// 国家/区号码。 例如：0086：中国大陆区号码。
	Areacode string `json:"areacode"`

	// 手机号码。
	Mobile string `json:"mobile"`

	// 验证码。 请调用“发送验证码”接口获取。
	VerificationCode string `json:"verification_code"`
}

func (o BankCardInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BankCardInfoV2 struct{}"
	}

	return strings.Join([]string{"BankCardInfoV2", string(data)}, " ")
}
