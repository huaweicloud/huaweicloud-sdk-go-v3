package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigTemplateRuleSetInfo struct {

	// **参数解释：** 语言类型。 **取值范围：** 自定义语言或者系统内置语言，当前支持的系统内置语言有： - ARKTS - C# - CANGJIE - C++ - CSS - GO - HTML - JAVA - JAVASCRIPT - KOTLIN - LUA - PHP - PYTHON - RUST - SCALA - SHELL - SQL - TERRAFORM - TYPESCRIPT - AZURERESOURCEMANAGER - DOCKER - KUBERNETES
	Language *string `json:"language,omitempty"`

	// **参数解释：** 规则集ID。 **取值范围：** 长度为[1~32]个字符。
	RuleSetId *string `json:"ruleSetId,omitempty"`
}

func (o ConfigTemplateRuleSetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigTemplateRuleSetInfo struct{}"
	}

	return strings.Join([]string{"ConfigTemplateRuleSetInfo", string(data)}, " ")
}
