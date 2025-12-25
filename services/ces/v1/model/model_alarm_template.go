package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmTemplate 创建的自定义告警模板详细信息
type AlarmTemplate struct {

	// **参数解释**： 自定义告警模板名称 **约束限制**： 不涉及。 **取值范围**： 以字母或汉字开头，可包含字母、数字、汉字、_、-，长度为[1,256]个字符，如：alarmTemplate-Test01。        **默认取值**： 不涉及。
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释**： 自定义告警模板描述。    **约束限制**： 不涉及。 **取值范围**： 长度范围[0,256]。          **默认取值**： 空字符串。
	TemplateDescription *string `json:"template_description,omitempty"`

	// **参数解释** 创建自定义告警模板选择的资源类型，即服务命名空间，如：选择弹性云服务器，则命名空间为SYS.ECS；各服务的命名空间可查看：“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释** 资源类型对应的指标监控维度，选择弹性云服务器，则维度为云服务器，dimension_name值为instance_id；各服务资源的指标维度名称可查看：“[服务指标维度](ces_03_0059.xml)”。 **约束限制**： 不涉及 **取值范围**： 必须以字母开头，多维度用\",\"分隔，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32。目前最大支持4个维度。字符串总最大长度为131。举例：单维度场景：instance_id；多维度场景：instance_id,disk        **默认取值**： 不涉及
	DimensionName *string `json:"dimension_name,omitempty"`

	// 自定义告警模板添加的一组或者多个告警策略。
	TemplateItems *[]TemplateItem `json:"template_items,omitempty"`

	// **参数解释** 自定义告警模版的ID，如：at1603330892378wkDm77y6B **约束限制**： 不涉及 **取值范围**： 以at开头，后跟字母、数字，长度最长为64 **默认取值**： 不涉及
	TemplateId *string `json:"template_id,omitempty"`
}

func (o AlarmTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplate struct{}"
	}

	return strings.Join([]string{"AlarmTemplate", string(data)}, " ")
}
