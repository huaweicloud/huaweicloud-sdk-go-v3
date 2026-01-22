package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAclRuleActionsResponse Response Object
type BatchUpdateAclRuleActionsResponse struct {

	// **参数解释**： 批量更新acl规则ID，为请求体中传入的规则ID **取值范围**： 不涉及
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchUpdateAclRuleActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAclRuleActionsResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateAclRuleActionsResponse", string(data)}, " ")
}
