package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePermRulesRequestBody 创建权限请求体
type CreatePermRulesRequestBody struct {

	// 权限信息，一次最多允许添加5条规则
	Rules []OnePermRuleRequestInfo `json:"rules"`
}

func (o CreatePermRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePermRulesRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePermRulesRequestBody", string(data)}, " ")
}
