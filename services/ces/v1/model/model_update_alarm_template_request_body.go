package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmTemplateRequestBody 修改自定义告警模板，请求参数
type UpdateAlarmTemplateRequestBody struct {

	// **参数解释** 告自定义告警模板名称 **约束限制**： 不涉及 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度为1-128。 **默认取值**： 不涉及
	TemplateName string `json:"template_name"`

	// **参数解释** 自定义告警模板详细描述 **约束限制**： 不涉及 **取值范围**： 长度范围[0,256]。 **默认取值**： 不涉及
	TemplateDescription *string `json:"template_description,omitempty"`

	// **参数解释** 创建自定义告警模板选择的资源类型，即服务命名空间，如：选择弹性云服务器，则命名空间为SYS.ECS；各服务的命名空间可查看：“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及
	Namespace string `json:"namespace"`

	// **参数解释** 资源类型对应的指标监控维度，选择弹性云服务器，则维度为云服务器，dimension_name值为instance_id；各服务资源的指标维度名称可查看：“[服务指标维度](ces_03_0059.xml)”。 **约束限制**： 不涉及 **取值范围**： 必须以字母开头，多维度用\",\"分隔，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32。目前最大支持4个维度。字符串总最大长度为131。举例：单维度场景：instance_id；多维度场景：instance_id,disk        **默认取值**： 不涉及
	DimensionName string `json:"dimension_name"`

	// 创建自定义告警模板添加一个或者多个指标的告警规则；目前最多可增加30组告警规则策略。
	TemplateItems []TemplateItem `json:"template_items"`
}

func (o UpdateAlarmTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateRequestBody", string(data)}, " ")
}
