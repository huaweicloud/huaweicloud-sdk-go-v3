package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskInfo Information of task
type TaskInfo struct {

	// **参数解释**: 待办任务的ID **取值范围**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 待办任务的引擎任务Id **取值范围**: 不涉及.
	AopengineTaskId *string `json:"aopengine_task_id,omitempty"`

	// **参数解释**: 待办任务的名称 **取值范围**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 待办任务的项目Id **取值范围**: 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 待办任务的描述 **取值范围**: 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 待办的创建时间 **取值范围**: 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**: 待办的创建者ID **取值范围**: system
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释**: 待办的创建者名字 **取值范围**: 不涉及
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释**: 待办的更新时间 **取值范围**: 不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**: 待办的修改者ID **取值范围**: system
	ModifierId *string `json:"modifier_id,omitempty"`

	// **参数解释**: 待办的修改者名字 **取值范围**: 不涉及
	ModifierName *string `json:"modifier_name,omitempty"`

	// **参数解释**: 待办支持的审核人用户ID **取值范围**: 不涉及
	ApproveuserId *string `json:"approveuser_id,omitempty"`

	// **参数解释**: 待办支持的审核人用户名字 **取值范围**: 不涉及
	ApproveuserName *string `json:"approveuser_name,omitempty"`

	// **参数解释**: 待办审核人用户名称 **取值范围**: 不涉及
	Approver *string `json:"approver,omitempty"`

	// **参数解释**: 待办的备注 **取值范围**: 不涉及
	Notes *string `json:"notes,omitempty"`

	// **参数解释**: 待办的节点流程拓扑图的KEY **取值范围**: 不涉及
	DefinitionKey *string `json:"definition_key,omitempty"`

	// **参数解释**: 待办的备注 **取值范围**: 不涉及
	Note *string `json:"note,omitempty"`

	// **参数解释**: 待办的超期时间 **取值范围**: 默认为创建时间+15天
	DueDate *string `json:"due_date,omitempty"`

	// **参数解释**: 待办的节点的流程或剧本ID 当 business_type是WORKFLOWPUBLISH或者WORKFLOWNODEAPPROVE，此时为流程ID 当 business_type是PLAYBOOKPUBLISH或者PLAYBOOKNODEAPPROVE，此时为剧本ID **取值范围**: 不涉及
	ActionId *string `json:"action_id,omitempty"`

	// **参数解释**: 待办的节点的流程或剧本版本ID 当 business_type是WORKFLOWPUBLISH或者WORKFLOWNODEAPPROVE，此时为流程版本ID 当 business_type是PLAYBOOKPUBLISH或者PLAYBOOKNODEAPPROVE，此时为剧本版本ID  **取值范围**: 不涉及
	ActionVersionId *string `json:"action_version_id,omitempty"`

	// **参数解释**: 待办的节点的流程或剧本的实例ID 当 business_type是WORKFLOWNODEAPPROVE，此时为流程实例ID 当 business_type是PLAYBOOKNODEAPPROVE，此时为剧本实例ID  **取值范围**: 不涉及
	ActionInstanceId *string `json:"action_instance_id,omitempty"`

	// **参数解释**: 待办任务的空间ID **取值范围**: 不涉及
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**: 待办任务审核意见 **取值范围**: 不涉及
	ReviewComments *string `json:"review_comments,omitempty"`

	// **参数解释**: 待办任务查看参数 **取值范围**: 不涉及
	ViewParameters *string `json:"view_parameters,omitempty"`

	// **参数解释**: 待办任务的人工处理参数 **取值范围**: 不涉及
	HandleParameters *string `json:"handle_parameters,omitempty"`

	// **参数解释**: 待办的业务类型 - WORKFLOWPUBLISH 流程发布 - WORKFLOWNODEAPPROVE 流程节点审核 - PLAYBOOKPUBLISH 剧本发布 - PLAYBOOKNODEAPPROVE 剧本节点审核  **取值范围**: - WORKFLOWPUBLISH - WORKFLOWNODEAPPROVE - PLAYBOOKPUBLISH - PLAYBOOKNODEAPPROVE
	BusinessType *string `json:"business_type,omitempty"`

	// **参数解释**: 待办任务的关联的流程 or 剧本名称 **取值范围**: 不涉及
	RelatedObject *string `json:"related_object,omitempty"`

	// **参数解释**: 待办节点的附件ID列表 **取值范围**: 不涉及
	AttachmentIdList *[]string `json:"attachment_id_list,omitempty"`

	// **参数解释**: 待办节点的待办评论 **取值范围**: 不涉及
	Comments *[]TaskCommentInfo `json:"comments,omitempty"`

	// **参数解释**: 待办的业务类型 - waiting  待处理 - canceled 已取消 - finished 已完成  **取值范围**: - waiting - canceled - finished
	Status *string `json:"status,omitempty"`

	// **参数解释**: 待办的到期处理方式 - CONTINUE 继续执行 - INTERRUPT 终止  **取值范围**: - CONTINUE - INTERRUPT
	DueHandle *string `json:"due_handle,omitempty"`
}

func (o TaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfo struct{}"
	}

	return strings.Join([]string{"TaskInfo", string(data)}, " ")
}
