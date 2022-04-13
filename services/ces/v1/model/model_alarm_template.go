package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建的自定义告警模板详细信息
type AlarmTemplate struct {
	// 自定义告警模板名称，如：alarmTemplate-Test01。

	TemplateName *string `json:"template_name,omitempty"`
	// 自定义告警模板描述。

	TemplateDescription *string `json:"template_description,omitempty"`
	// 自定义告警模板选择的资源类型，即服务命名空间，如：选择弹性云服务器，则命名空间为SYS.ECS，各资源的监控指标名称可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Namespace *string `json:"namespace,omitempty"`
	// 自定义告警模板选择的资源维度，如：弹性云服务器，则维度为instance_id，各服务资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	DimensionName *string `json:"dimension_name,omitempty"`
	// 自定义告警模板添加的一组或者多个告警策略。

	TemplateItems *[]TemplateItem `json:"template_items,omitempty"`
	// 自定义告警模板的ID，如：at1603330892378wkDm77y6B。

	TemplateId *string `json:"template_id,omitempty"`
}

func (o AlarmTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplate struct{}"
	}

	return strings.Join([]string{"AlarmTemplate", string(data)}, " ")
}
