package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertRevokeConfigRequestBody struct {

	// **参数解释**： 用户ID，获取方式参见本接口的接口约束 **约束限制**： 与agencyId互斥，二选一填写 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserId *string `json:"userId,omitempty"`

	// **参数解释**： 用户ID，获取方式参见[如何获取用户ID](cce_02_0249.xml#section1) **约束限制**： 与agencyId互斥，二选一填写 **取值范围**： 不涉及 **默认取值**： 不涉及
	AgencyId *string `json:"agencyId,omitempty"`
}

func (o CertRevokeConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertRevokeConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CertRevokeConfigRequestBody", string(data)}, " ")
}
