package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskRulesetInfo struct {

	// **参数解释：** 规则集ID。 **取值范围:** 长度为[1~32]个字符。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释：** 规则集语言。 **取值范围:** 自定义语言或者系统内置语言，当前支持的系统内置语言有： - ARKTS - C# - CANGJIE - C++ - CSS - GO - HTML - JAVA - JAVASCRIPT - KOTLIN - LUA - PHP - PYTHON - RUST - SCALA - SHELL - SQL - TERRAFORM - TYPESCRIPT - AZURERESOURCEMANAGER - DOCKER - KUBERNETES
	Language *string `json:"language,omitempty"`

	// **参数解释：** 规则集名称。 **取值范围:** 任意中英文，数字，点，连字符-及下划线_。长度为[1~128]个字符。
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释：** 规则集选取状态。 **取值范围:** - selected：已选。 - optional：可选，指未开启的语言在当前项目下的默认规则集。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 规则集属性。 **取值范围:** - 0：用户自定义规则集。 - 1：系统规则集。
	Status *string `json:"status,omitempty"`
}

func (o TaskRulesetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskRulesetInfo struct{}"
	}

	return strings.Join([]string{"TaskRulesetInfo", string(data)}, " ")
}
