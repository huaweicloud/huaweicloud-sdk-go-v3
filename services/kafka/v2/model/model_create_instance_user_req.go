package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstanceUserReq struct {
	// 用户名称。

	UserName *string `json:"user_name,omitempty"`
	// 用户密码。  密码不能和用户名相同。 复杂度要求： - 输入长度为8到32位的字符串。 - 必须包含如下四种字符中的两种组合：   - 小写字母   - 大写字母   - 数字   - 特殊字符包括（`~!@#$%^&*()-_=+\\|[{}]:'\",<.>/?）

	UserPasswd *string `json:"user_passwd,omitempty"`
}

func (o CreateInstanceUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceUserReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceUserReq", string(data)}, " ")
}
