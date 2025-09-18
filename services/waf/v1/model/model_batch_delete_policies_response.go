package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoliciesResponse Response Object
type BatchDeletePoliciesResponse struct {

	// **参数解释：** 策略id数组，返回已成功批量删除的防护策略唯一标识。 **约束限制：** 不涉及 **取值范围：** 数组内元素为字符串类型，与请求体中传入的有效策略ID对应 **默认取值：** 不涉及
	PolicyIds      *[]string `json:"policy_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeletePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePoliciesResponse", string(data)}, " ")
}
