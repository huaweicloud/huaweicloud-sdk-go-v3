package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Credentials 委托凭据信息
type Credentials struct {

	// **参数解释**： 临时安全凭证的AK **约束限制**： 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AccessKeyId *string `json:"accessKeyId,omitempty"`

	// **参数解释：** 临时安全凭证的SK **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`

	// **参数解释**： 临时安全凭证的security_token **约束限制**： 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SecurityToken *string `json:"securityToken,omitempty"`

	// **参数解释：** 临时安全凭证的失效时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Expiration *string `json:"expiration,omitempty"`
}

func (o Credentials) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Credentials struct{}"
	}

	return strings.Join([]string{"Credentials", string(data)}, " ")
}
