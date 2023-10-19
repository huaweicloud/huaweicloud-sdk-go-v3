package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteAclRulesResponseData struct {

	// 批量删除规则返回data
	ResponseDatas *[]BatchDeleteRuleInfo `json:"responseDatas,omitempty"`
}

func (o BatchDeleteAclRulesResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAclRulesResponseData struct{}"
	}

	return strings.Join([]string{"BatchDeleteAclRulesResponseData", string(data)}, " ")
}
