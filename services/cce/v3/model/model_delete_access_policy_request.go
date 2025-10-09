package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessPolicyRequest Request Object
type DeleteAccessPolicyRequest struct {

	// **参数解释：** 访问策略ID。获取方式请参见[获取访问策略列表](ListAccessPolicy.xml)。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`
}

func (o DeleteAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAccessPolicyRequest", string(data)}, " ")
}
