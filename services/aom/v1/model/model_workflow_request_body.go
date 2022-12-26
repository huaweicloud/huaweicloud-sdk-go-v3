package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自动化运维工作流工作流元数据。
type WorkflowRequestBody struct {

	// 工作流名称，需要满足中文、英文大小写、数字、中划线和下划线{1,64}。
	Name *string `json:"name,omitempty"`

	// 工作流类型，可以为cron、manual
	Type *string `json:"type,omitempty"`

	// 工作流描述信息。
	Description *string `json:"description,omitempty"`

	// 标签键和值列表，标签键值对数量范围是0至20。
	Tags *interface{} `json:"tags,omitempty"`

	// 模板名称，示例：CMS::ECS::BulkyRunScript  CMS::ECS::BulkyStartECSInstances CMS::ECS::BulkyCleanDisks
	TemplateName *string `json:"template_name,omitempty"`

	// 模板id。
	TemplateId *string `json:"template_id,omitempty"`

	// 任务执行时需要的参数列表。
	Input map[string]interface{} `json:"input,omitempty"`

	// 引用，参数引用。
	Quote *[]string `json:"quote,omitempty"`

	Trigger *Trigger `json:"trigger,omitempty"`

	// 作业名称。
	JobName *string `json:"job_name,omitempty"`

	// 作业id。
	JobId *string `json:"job_id,omitempty"`

	// 服务场景分类。
	ServiceScenario *string `json:"service_scenario,omitempty"`

	// 服务名称。
	ServiceName *string `json:"service_name,omitempty"`

	// 任务类型。package,script,job,cloud,standard,customize
	TaskType *string `json:"task_type,omitempty"`
}

func (o WorkflowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowRequestBody struct{}"
	}

	return strings.Join([]string{"WorkflowRequestBody", string(data)}, " ")
}
