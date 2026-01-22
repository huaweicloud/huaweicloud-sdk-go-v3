package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRuleInfo struct {

	// **参数解释**： 批量删除的ACL的名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 批量删除的ACL的ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o BatchDeleteRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRuleInfo struct{}"
	}

	return strings.Join([]string{"BatchDeleteRuleInfo", string(data)}, " ")
}
