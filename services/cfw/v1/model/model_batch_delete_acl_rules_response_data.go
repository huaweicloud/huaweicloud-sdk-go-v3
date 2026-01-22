package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteAclRulesResponseData struct {

	// **参数解释**： 批量删除规则返回 **取值范围**： 不涉及
	ResponseDatas *[]BatchDeleteRuleInfo `json:"responseDatas,omitempty"`
}

func (o BatchDeleteAclRulesResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAclRulesResponseData struct{}"
	}

	return strings.Join([]string{"BatchDeleteAclRulesResponseData", string(data)}, " ")
}
