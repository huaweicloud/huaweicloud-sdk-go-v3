package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyDescription **参数解释**： 身份策略描述信息。  **约束限制**： 不涉及。  **取值范围**： 长度范围为[0,1000]。  **默认取值**： 不涉及。
type PolicyDescription struct {
}

func (o PolicyDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDescription struct{}"
	}

	return strings.Join([]string{"PolicyDescription", string(data)}, " ")
}
