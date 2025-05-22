package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPasswordRequestBody **参数解释**： 重置密码的请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ResetPasswordRequestBody struct {

	// **参数解释**： GaussDB(DWS) 集群管理员新密码。 **约束限制**： 新密码复杂度要求如下：  - 密码字符长度为12~32位。 - 不能与用户名或倒序的用户名相同。 - 至少包含以下4种类型的3种：    - 小写字母   - 大写字母   - 数字   - 特殊字符（~!?,.:;-_'\"(){}[]/<>@#%^&*+|\\=）。 - 不能与历史密码相同。 - 不能为弱密码。  **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NewPassword string `json:"new_password"`
}

func (o ResetPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ResetPasswordRequestBody", string(data)}, " ")
}
