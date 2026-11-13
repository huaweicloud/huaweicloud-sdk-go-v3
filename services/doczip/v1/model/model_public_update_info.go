package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicUpdateInfo struct {

	// **参数解释**： Key ID。 **约束限制**： 不涉及。 **取值范围**： [1-128]个字符。 **默认取值**： 不涉及。
	Kid *string `json:"kid,omitempty"`

	// **参数解释**： PEM 格式公钥。 **约束限制**： 不涉及。 **取值范围**： [1-2048]的字符串。 **默认取值**： 不涉及。
	PublicKey *string `json:"public_key,omitempty"`

	// **参数解释**： 操作类型。 **约束限制**： 不涉及。 **取值范围**： - add：新增。 - update：更新。 - delete：删除 **默认取值**： 不涉及。
	OpType *string `json:"op_type,omitempty"`
}

func (o PublicUpdateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicUpdateInfo struct{}"
	}

	return strings.Join([]string{"PublicUpdateInfo", string(data)}, " ")
}
