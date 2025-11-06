package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmRuleTemplateBody struct {

	// 告警模板名称。
	AlarmRuleTemplateName string `json:"alarm_rule_template_name"`

	// 告警模板英文名称。
	AlarmRuleTemplateNameEn *string `json:"alarm_rule_template_name_en,omitempty"`

	// 告警模板描述。
	AlarmRuleTemplateDescription string `json:"alarm_rule_template_description"`

	// 告警模板id。
	AlarmRuleTemplateId string `json:"alarm_rule_template_id"`

	// 项目ID，可以从控制台获取，也可以从调用API处获取。获取方式请参见：[获取项目ID](aom_04_0024.xml)。
	AlarmRuleTemplateProjectId *string `json:"alarm_rule_template_project_id,omitempty"`

	// 告警模板类型 - “statics”：静态告警模板 - “dynamic”：动态告警模板
	AlarmRuleTemplateType string `json:"alarm_rule_template_type"`

	// 告警模板来源。
	AlarmRuleTemplateSource *string `json:"alarm_rule_template_source,omitempty"`

	// 告警模板所绑定的告警规则。
	AlarmRuleTemplateBinding map[string]string `json:"alarm_rule_template_binding"`

	// 告警模板列表。
	AlarmTemplateSpecList []AlarmRuleTemplateSpecWithCloudService `json:"alarm_template_spec_list"`

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间。
	ModifyTime *int64 `json:"modify_time,omitempty"`

	Templating *Templating `json:"templating"`

	// 告警模板版本。
	TemplateVersion string `json:"template_version"`
}

func (o AlarmRuleTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRuleTemplateBody struct{}"
	}

	return strings.Join([]string{"AlarmRuleTemplateBody", string(data)}, " ")
}
