package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建自定义告警模板请求参数。
type CreateAlarmTemplateRequestBody struct {
	// 自定义告警模板名称，只能包含0-9/a-z/A-Z/_/-或汉字，长度为1-128。。

	TemplateName string `json:"template_name"`
	// 自定义告警模板详细描述，长度为0-256。

	TemplateDescription *string `json:"template_description,omitempty"`
	// 创建自定义告警模板选择的资源类型，即服务命名空间，如：选择弹性云服务器，则命名空间为SYS.ECS；各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Namespace string `json:"namespace"`
	// 资源类型对应的指标监控维度，选择弹性云服务器，则维度为云服务器，dimension_name值为instance_id；各服务资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

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
