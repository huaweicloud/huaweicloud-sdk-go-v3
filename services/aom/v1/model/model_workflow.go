package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Workflow 变更服务工作流工作流元数据。
type Workflow struct {

	// 工作流id，唯一标识，根据project_id和workflow_name生成。
	Id *string `json:"id,omitempty"`

	// 工作流名称。
	Name string `json:"name"`

	// 工作流类型，可以为cron、manual
	Type *string `json:"type,omitempty"`

	// 工作流描述信息。
	Description *string `json:"description,omitempty"`

	// 标签键和值列表，标签键值对数量范围是0至20。
	Tags map[string]string `json:"tags,omitempty"`

	// 工作流创建时间，为utc时间毫秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 工作流创人，从接口调用传入的token中获取。
	CreateBy *string `json:"create_by,omitempty"`

	// 工作流更新时间，为utc时间毫秒数。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 工作流更新人，从接口调用传入的token中获取。
	UpdateBy *string `json:"update_by,omitempty"`

	// 模板名称。
	TemplateName *string `json:"template_name,omitempty"`

	// 模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 任务执行时需要的参数列表。
	Input map[string]interface{} `json:"input,omitempty"`

	// 最近一次执行id，也是工作流id
	LastExecutionId *string `json:"last_execution_id,omitempty"`

	// 任务状态，包含success，fail,executing
	Status *string `json:"status,omitempty"`

	// 工作流的引用。
	CitationUrns *[]string `json:"citation_urns,omitempty"`

	// 最近一次执行结束时间，为utc时间毫秒数
	LastExecutionEndTime *int64 `json:"last_execution_end_time,omitempty"`

	// 最近一次执行开始时间，为utc时间毫秒数
	LastExecutionStartTime *int64 `json:"last_execution_start_time,omitempty"`

	// 引用，参数引用
	Quote *[]string `json:"quote,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 作业id
	JobId *string `json:"job_id,omitempty"`

	// 服务场景分类
	ServiceScenario *string `json:"service_scenario,omitempty"`

	// 服务名称
	ServiceName *string `json:"service_name,omitempty"`

	// 任务类型
	TaskType *string `json:"task_type,omitempty"`

	// functiongraph返回的PROJECT_ID
	ProjectId *string `json:"project_id,omitempty"`

	// functiongraph返回的WORKFLOW_ID
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 任务状态
	TaskStatus *string `json:"task_status,omitempty"`

	// 任务节点
	Nodes *[]Node `json:"nodes,omitempty"`

	// 编辑时间
	EditTime *int64 `json:"edit_time,omitempty"`

	// 执行动作细粒度权限
	ExecutionActionRules *[]string `json:"execution_action_rules,omitempty"`

	// 云服务权限
	ExecutionPermission *[]string `json:"execution_permission,omitempty"`

	// 全局参数
	GlobalParameters *[]Parameter `json:"global_parameters,omitempty"`

	// 逻辑删除
	IsDelete *bool `json:"is_delete,omitempty"`

	// 任务步骤
	Steps *[]Step `json:"steps,omitempty"`

	// 任务输出
	Output *string `json:"output,omitempty"`

	// 触发器id
	TriggerId *string `json:"trigger_id,omitempty"`

	// 触发器状态
	TriggerStatus *string `json:"trigger_status,omitempty"`

	// 审批id
	ApproveId *string `json:"approve_id,omitempty"`

	TemplateI18n *WorkFlowModel `json:"template_i18n,omitempty"`

	// 任务所属的企业项目
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务最后一次执行人
	LastExecuteBy *string `json:"last_execute_by,omitempty"`
}

func (o Workflow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workflow struct{}"
	}

	return strings.Join([]string{"Workflow", string(data)}, " ")
}
