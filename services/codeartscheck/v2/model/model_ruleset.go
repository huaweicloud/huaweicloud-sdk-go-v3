package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Ruleset 创建规则集信息。
type Ruleset struct {

	// **参数解释:** 项目的32位UUID，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/zh-cn/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)接口获取。 **约束限制:** 不涉及。 **取值范围:** 长度为[1~32]个字符。 **默认取值:** 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释:** 新规则集名称。 **约束限制:** 不能含有任意中英文，数字，点，连字符-及下划线_以外的特殊字符。 **取值范围:** 任意中英文，数字，点，连字符-及下划线_。长度为[1~128]个字符。 **默认取值:** 不涉及。
	TemplateName string `json:"template_name"`

	// **参数解释:** 规则集语言。 **约束限制:** 只能选择当前系统上支持的语言。 **取值范围:** 只能传一种语言。当前支持的语言有： - ARKTS - C# - CANGJIE - C++ - CSS - GO - HTML - JAVA - JAVASCRIPT - KOTLIN - LUA - PHP - PYTHON - RUST - SCALA - SHELL - SQL - TERRAFORM - TYPESCRIPT - AZURERESOURCEMANAGER - DOCKER - KUBERNETES **默认取值:** 不涉及。
	Language string `json:"language"`

	// **参数解释:** 是否基于已存在的规则集创建新规则集。 **约束限制:** 只能填0或1。 **取值范围:** - 0：新建空白规则集。 - 1：基于已存在的规则集创建新规则集。 **默认取值:** 不涉及。
	IsDefault string `json:"is_default"`

	// **参数解释:** 新启用规则ID集合，启用多条规则时，id用英文逗号相连。可调用[[获取规则列表接口](https://support.huaweicloud.com/api-codecheck/ListRules.html)](tag:hws)[[获取规则列表接口](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ListRules.html)](tag:hws_hk)[获取规则列表接口](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)获取规则ID。 **约束限制:** 规则ID为32位UUID，多个ID之间用英文逗号分隔。例如：\"00034fc32ee045388365d42ca4e07fe1,002c2eda6d6c11edab16fa163e0fa374\"。 **取值范围:** 长度为[1~4000]个字符。 **默认取值:** 不涉及。
	RuleIds *string `json:"rule_ids,omitempty"`

	// **参数解释:** 新关闭规则ID集合，多个ID用英文逗号相连。可调用[[获取规则列表接口](https://support.huaweicloud.com/api-codecheck/ListRules.html)](tag:hws)[[获取规则列表接口](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ListRules.html)](tag:hws_hk)[获取规则列表接口](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)获取规则ID。 **约束限制:** 规则ID为32位UUID，多个ID之间用英文逗号分隔。例如：\"00034fc32ee045388365d42ca4e07fe1,002c2eda6d6c11edab16fa163e0fa374\"。 **取值范围:** 长度为[1~4000]个字符。 **默认取值:** 不涉及。
	UncheckIds *string `json:"uncheck_ids,omitempty"`

	// **参数解释:** 规则集ID。当is_default参数为1时，该参数需要传参。可调用[[查询规则集列表](https://support.huaweicloud.com/api-codecheck/ListRulesets.html)](tag:hws)[[查询规则集列表](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ListRulesets.html)](tag:hws_hk)[查询规则集列表](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)获取已存在的规则集ID。 **约束限制:** 32位UUID。例如：\"da303d80d2594faca60dbcc2e29c18aa\"。 **取值范围:** 不涉及。 **默认取值:** 不涉及。
	TemplateId *string `json:"template_id,omitempty"`

	// 自定义规则参数项，支持修改规则阈值。
	CustomAttributes *[]CustomAttributes `json:"custom_attributes,omitempty"`
}

func (o Ruleset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ruleset struct{}"
	}

	return strings.Join([]string{"Ruleset", string(data)}, " ")
}
