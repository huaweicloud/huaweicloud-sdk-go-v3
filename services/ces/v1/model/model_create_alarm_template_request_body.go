package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlarmTemplateRequestBody 创建自定义告警模板请求参数。
type CreateAlarmTemplateRequestBody struct {

	// **参数解释** 告自定义告警模板名称 **约束限制**： 不涉及 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度为1-128。 **默认取值**： 不涉及
	TemplateName string `json:"template_name"`

	// **参数解释** 自定义告警模板详细描述 **约束限制**： 不涉及 **取值范围**： 长度范围[0,256]。 **默认取值**： 不涉及
	TemplateDescription *string `json:"template_description,omitempty"`

	// **参数解释** 创建自定义告警模板选择的资源类型，即服务命名空间，如：选择弹性云服务器，则命名空间为SYS.ECS；各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。 **约束限制**： 不涉及 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及
	Namespace string `json:"namespace"`

	// **参数解释** 资源类型对应的指标监控维度，选择弹性云服务器，则维度为云服务器，dimension_name值为instance_id；各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **约束限制**： 不涉及 **取值范围**： 必须以字母开头，多维度用\",\"分隔，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32。目前最大支持4个维度。字符串总最大长度为131。举例：单维度场景：instance_id；多维度场景：instance_id,disk        **默认取值**： 不涉及
	DimensionName string `json:"dimension_name"`

	// 创建自定义告警模板添加一个或者多个指标的告警规则；目前最多可增加30组告警规则策略。
	TemplateItems []TemplateItem `json:"template_items"`
}

func (o CreateAlarmTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateRequestBody", string(data)}, " ")
}
