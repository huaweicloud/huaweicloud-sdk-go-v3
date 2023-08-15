package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PwdResetRequest 数据库root用户密码。  8~32个字符。 至少包含以下字符中的三种： 大写字母、小写字母、数字和特殊字符~!@#%^*-_=+?,  弱密码校验。
type PwdResetRequest struct {

	// 数据库root用户密码。  - 8~32个字符。 - 至少包含以下字符中的三种： - 大写字母、小写字母、数字和特殊字符~!@#%^*-_=+?,  - 弱密码校验。
	Password string `json:"password"`
}

func (o PwdResetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PwdResetRequest struct{}"
	}

	return strings.Join([]string{"PwdResetRequest", string(data)}, " ")
}
