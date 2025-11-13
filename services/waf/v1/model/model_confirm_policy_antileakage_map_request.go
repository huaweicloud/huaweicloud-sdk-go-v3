package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmPolicyAntileakageMapRequest Request Object
type ConfirmPolicyAntileakageMapRequest struct {

	// **参数解释：** 语言的类型 - cn代表中文 - en代表英文  **约束限制：** 不涉及 **取值范围：** - cn - en  **默认取值：** - cn
	Lang string `json:"lang"`
}

func (o ConfirmPolicyAntileakageMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmPolicyAntileakageMapRequest struct{}"
	}

	return strings.Join([]string{"ConfirmPolicyAntileakageMapRequest", string(data)}, " ")
}
