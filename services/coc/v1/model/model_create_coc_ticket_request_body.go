package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCocTicketRequestBody struct {

	// **参数解释：** 工单标题。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Title *string `json:"title,omitempty"`

	// **参数解释：** 变更单描述。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChangeNotes *string `json:"change_notes,omitempty"`

	// **参数解释：** 问题单描述信息。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 企业项目ID **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段必填。 **取值范围：** 不涉及 **默认取值：** 0
	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	// **参数解释：** 变更类型。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。 **取值范围：** 枚举值 change_type_conventional：常规变更 change_type_urgentu：紧急变更 **默认取值：** 不涉及
	ChangeType *string `json:"change_type,omitempty"`

	// **参数解释：** 工单级别。 **约束限制：** 不涉及 **取值范围：** 当ticket_type为change创建变更单时，枚举值 change_level_010 -- A级变更 change_level_020 -- B级变更 change_level_030 -- C级变更 change_level_040 -- D级变更 当ticket_type为issues_mgmt创建问题单时，枚举值 issues_level_1000 -- 致命 issues_level_2000 -- 严重 issues_level_3000 -- 一般 issues_level_4000 -- 提示 **默认取值：** 不涉及
	Level *string `json:"level,omitempty"`

	// **参数解释：** 问题单类型，通过访问 云运维中心-->基础配置-->流程管理页签下问题流程-->问题类别可查询所有可传递的问题类型，此处传递问题类型ID。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	TicketType *string `json:"ticket_type,omitempty"`

	// **参数解释：** 任务类型，可选作业或者变更指导书两种。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。当取值为change_scheme_guides时，请求参数change_guides必填。当取值为change_scheme_runbook时，参数plan_task_sub_type、plan_task_id、plan_task_name和plan_task_params必填。 **取值范围：** 枚举值 change_scheme_runbook：作业 change_scheme_guides：变更指导书 **默认取值：** 不涉及
	ChangeScheme *string `json:"change_scheme,omitempty"`

	// **参数解释：** 变更指导书ID。 **约束限制：** 当ticket_type为change创建变更单，且任务选择变更指导书时，该字段必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChangeGuides *string `json:"change_guides,omitempty"`

	// **参数解释：** 问题单附件，上传附件ID。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	CommitUploadAttachment *string `json:"commit_upload_attachment,omitempty"`

	// **参数解释：** 问题单所属Region，此处传RegionID，多个Region用英文逗号隔开。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Regions *string `json:"regions,omitempty"`

	// **参数解释：** 变更场景。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。 **取值范围：** 取配置页面【流程管理】下“变更场景”TAB页中列表属性ID列的值，示例：GOCMLL06 **默认取值：** 不涉及
	ChangeSceneId *string `json:"change_scene_id,omitempty"`

	// **参数解释：** 归属应用。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	CurrentCloudServiceId *string `json:"current_cloud_service_id,omitempty"`

	// **参数解释：** 归属应用。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	RootCauseCloudService *string `json:"root_cause_cloud_service,omitempty"`

	// **参数解释：** 问题单来源。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。 **取值范围：** 当ticket_type为issues_mgmt创建问题单时，枚举值 issues_mgmt_associated_type_1000 -- 事件 issues_mgmt_associated_type_4000 -- 运维主动发现 issues_mgmt_associated_type_2000 -- 告警 issues_mgmt_associated_type_3000 -- WarRoom **默认取值：** 不涉及
	Source *string `json:"source,omitempty"`

	// **参数解释：** 问题单来源工单单号。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。当source的值为issues_mgmt_associated_type_1000、issues_mgmt_associated_type_2000或issues_mgmt_associated_type_3000时，必填。需要填写关联的工单单号。 **取值范围：** 不涉及 **默认取值：** 不涉及
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释：** 发现时间，时间戳。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	FountTime *int64 `json:"fount_time,omitempty"`

	// **参数解释：** 问题单处理人类型。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段必填。 **取值范围：** 当ticket_type为issues_mgmt创建问题单时，枚举值 issues_mgmt_virtual_schedule_type_1000 -- 排班 issues_mgmt_virtual_schedule_type_2000 -- 个人 **默认取值：** 不涉及
	VirtualScheduleType *string `json:"virtual_schedule_type,omitempty"`

	// **参数解释：** 问题单责任人工号ID。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。如需指定问题单责任人，则该字段必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	IssueContactPerson *string `json:"issue_contact_person,omitempty"`

	// **参数解释：** 问题单责任人从排班中获取，该值为排班场景ID。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	ScheduleScenes *string `json:"schedule_scenes,omitempty"`

	// **参数解释：** 问题单责任人从排班中获取，该值为排班角色ID。 **约束限制：** 当ticket_type为issues_mgmt创建问题单时，该字段选填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	VirtualScheduleRole *string `json:"virtual_schedule_role,omitempty"`

	// **参数解释：** 变更区域ID。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	LocationId *string `json:"location_id,omitempty"`

	// **参数解释：** 预案子类型。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。当任务类型change_scheme取值为change_scheme_runbook时，该值必填。 **取值范围：** 枚举值 CUSTOMIZATION：自定义作业 COMMUNAL：公共作业 **默认取值：** 不涉及
	PlanTaskSubType *string `json:"plan_task_sub_type,omitempty"`

	// **参数解释：** 需要执行的作业ID。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。当任务类型change_scheme取值为change_scheme_runbook时，该值必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	PlanTaskId *string `json:"plan_task_id,omitempty"`

	// **参数解释：** 需要执行的作业名称。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。当任务类型change_scheme取值为change_scheme_runbook时，该值必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	PlanTaskName *string `json:"plan_task_name,omitempty"`

	// **参数解释：** 执行作业时所需的参数信息。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填。当任务类型change_scheme取值为change_scheme_runbook时，该值必填。 **取值范围：** 不涉及 **默认取值：** 不涉及
	PlanTaskParams *string `json:"plan_task_params,omitempty"`

	// **参数解释：** 是否启动流程，当此值为false时，创建出来的工单为草稿状态。此值默认为true，创建出来的工单状态为未受理状态。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IsStartProcess *bool `json:"is_start_process,omitempty"`

	// **参数解释：** 变更子单的信息，包含变更单涉及的服务和Region信息。 **约束限制：** 当ticket_type为change创建变更单时，该字段必填且有效，当ticket_type非change时，该字段可置空。 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubTickets *[]TicketCreateSubTicketInfo `json:"sub_tickets,omitempty"`
}

func (o CreateCocTicketRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCocTicketRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCocTicketRequestBody", string(data)}, " ")
}
