package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowResponse Response Object
type UpdateWorkflowResponse struct {

	// Workflow工作流名称，1到64位只包含中英文、数字、空格、下划线（_）和中划线（-），并且以中英文开头。
	Name *string `json:"name,omitempty"`

	// Workflow工作流ID。创建工作流时后台自动生成。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// Workflow工作流的创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// Workflow工作流的描述信息。
	Description *string `json:"description,omitempty"`

	// Workflow工作流包含的步骤定义。
	Steps *[]WorkflowStep `json:"steps,omitempty"`

	// 创建Workflow工作流的用户名。
	UserName *string `json:"user_name,omitempty"`

	// 工作空间ID。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Workflow需要的数据。
	DataRequirements *[]DataRequirement `json:"data_requirements,omitempty"`

	// Workflow包含的数据。
	Data *[]Data `json:"data,omitempty"`

	// Workflow包含的参数。
	Parameters *[]WorkflowParameter `json:"parameters,omitempty"`

	// 从指定Workflow工作流进行复制。通过复制来创建Workflow时必填。
	SourceWorkflowId *string `json:"source_workflow_id,omitempty"`

	GallerySubscription *WorkflowGallerySubscription `json:"gallery_subscription,omitempty"`

	LatestExecution *ExecutionBrief `json:"latest_execution,omitempty"`

	// 工作流的已运行次数。
	RunCount *int32 `json:"run_count,omitempty"`

	// 当前工作流的必选参数是否都已填完。
	ParamReady *bool `json:"param_ready,omitempty"`

	// 工作流来源，可选值为ai_gallery，表示工作流是从AI Gallery导入的。
	Source *string `json:"source,omitempty"`

	// Workflow包含的统一存储定义。
	Storages *[]WorkflowStorage `json:"storages,omitempty"`

	// 为Workflow工作流设置的标签。
	Labels *[]string `json:"labels,omitempty"`

	// 工作流绑定的资产。
	Assets *[]WorkflowAsset `json:"assets,omitempty"`

	// 工作流包含的子图。
	SubGraphs *[]WorkflowSubgraph `json:"sub_graphs,omitempty"`

	// 计费工作流使用的拓展字段。
	Extend map[string]interface{} `json:"extend,omitempty"`

	Policy *WorkflowPolicy `json:"policy,omitempty"`

	// 工作流SMN消息订阅开关，默认为false，表示关闭消息订阅开关。
	WithSubscription *bool `json:"with_subscription,omitempty"`

	// SMN开关。
	SmnSwitch *bool `json:"smn_switch,omitempty"`

	// SMN消息订阅ID。
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// 自动学习模板ID。
	ExemlTemplateId *string `json:"exeml_template_id,omitempty"`

	// 最近一次修改的时间。
	LastModifiedAt *string `json:"last_modified_at,omitempty"`

	Package        *WorkflowServicePackege `json:"package,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowResponse", string(data)}, " ")
}
