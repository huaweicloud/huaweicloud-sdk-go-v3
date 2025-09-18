package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeletePoliciesRequestBody struct {

	// **参数解释：** 策略id数组，包含待批量删除的防护策略唯一标识，从查询防护策略列表接口获取。 **约束限制：** 不涉及 **取值范围：** 数组内元素为字符串类型，每个元素对应一个防护策略ID **默认取值：** 不涉及
	PolicyIds *[]string `json:"policy_ids,omitempty"`
}

func (o BatchDeletePoliciesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoliciesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeletePoliciesRequestBody", string(data)}, " ")
}
