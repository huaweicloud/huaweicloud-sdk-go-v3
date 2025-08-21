package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessPolicy struct {

	// **参数解释**： 变化数量 **取值范围**： 不涉及
	Changed *int32 `json:"changed,omitempty"`

	// **参数解释**： EIP访问控制策略 **取值范围**： 不涉及
	Eip *int32 `json:"eip,omitempty"`

	// **参数解释**： NAT访问控制策略 **取值范围**： 不涉及
	Nat *int32 `json:"nat,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o AccessPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicy struct{}"
	}

	return strings.Join([]string{"AccessPolicy", string(data)}, " ")
}
