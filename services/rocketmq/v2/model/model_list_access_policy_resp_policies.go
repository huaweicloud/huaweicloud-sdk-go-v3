package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAccessPolicyRespPolicies struct {

	// **参数解释**： 用户名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AccessKey *string `json:"access_key,omitempty"`

	// **参数解释**： IP白名单。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	// **参数解释**： 是否为管理员。 **约束限制**： 不涉及。 **取值范围**： - true：是管理员。 - false：不是管理员。 **默认取值**： 不涉及。
	Admin *bool `json:"admin,omitempty"`

	// **参数解释**： 用户权限。 **约束限制**： 不涉及。 **取值范围**： - PUB：拥有发布权限。 - SUB：拥有订阅权限。 - PUB|SUB：拥有发布订阅权限。 - DENY：无权限。 **默认取值**： 不涉及。
	Perm *string `json:"perm,omitempty"`
}

func (o ListAccessPolicyRespPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPolicyRespPolicies struct{}"
	}

	return strings.Join([]string{"ListAccessPolicyRespPolicies", string(data)}, " ")
}
