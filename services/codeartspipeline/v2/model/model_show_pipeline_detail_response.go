package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineDetailResponse Response Object
type ShowPipelineDetailResponse struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 流水线名称。 **取值范围**： 仅包含中文、大小写英文字母、数字、'-'和'_'，且长度为[1,128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 对流水线的补充描述。 **取值范围**： 不超过1024字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 流水线版本，默认为3.0。 **取值范围**： 不涉及。
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// **参数解释**： 当前环境所属局点。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 所属租户ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 所属微服务ID。可以通过[查询微服务列表](ListMicroservice.xml)接口获取，其中data.id即为微服务ID。 **取值范围**： 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释**： 是否为变更流水线。 **取值范围**： - true：是变更流水线。 - false：不是变更流水线。
	IsPublish *bool `json:"is_publish,omitempty"`

	// **参数解释**： 流水线创建人ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释**： 流水线创建人名称。 **取值范围**： 不涉及。
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释**： 流水线上次更新人ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	UpdaterId *string `json:"updater_id,omitempty"`

	// **参数解释**： 流水线创建时间。 **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 流水线更新时间。 **取值范围**： 不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 流水线是否被当前用户收藏。 **取值范围**： - true：流水线已被收藏。 - false：流水线未被收藏。
	IsCollect *bool `json:"is_collect,omitempty"`

	// **参数解释**： 流水线源列表。 **取值范围**： 不涉及。
	Sources *[]PipelineSource `json:"sources,omitempty"`

	// **参数解释**： 流水线自定义参数。 **取值范围**： 不涉及。
	Variables *[]PipelineVariable `json:"variables,omitempty"`

	// **参数解释**： 流水线定时任务设置。 **取值范围**： 不涉及。
	Schedules *[]PipelineSchedule `json:"schedules,omitempty"`

	// **参数解释**： 流水线事件触发设置。 **取值范围**： 不涉及。
	Triggers *[]PipelineTrigger `json:"triggers,omitempty"`

	// **参数解释**： 流水线所属分组ID。 **取值范围**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 流水线定义JSON。 **取值范围**： 不涉及。
	Definition *string `json:"definition,omitempty"`

	// **参数解释**： 流水线涉密等级。 **取值范围**： 不涉及。
	SecurityLevel *int32 `json:"security_level,omitempty"`

	// **参数解释**： 复制流水线场景下，原流水线ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	OriginId *string `json:"origin_id,omitempty"`

	// **参数解释**： 是否禁用发布分支管理。 **取值范围**： - true：禁用发布分支管理。 - false：不禁用发布分支管理。
	DisableReleaseBranchManagement *bool `json:"disable_release_branch_management,omitempty"`

	// **参数解释**： 流水线是否已被删除。 **取值范围**： - true：已删除。 - false：未删除。
	Deleted *bool `json:"deleted,omitempty"`

	// **参数解释**： 流水线是否被禁用。 **取值范围**： - true：已禁用。 - false：未禁用。
	Banned *bool `json:"banned,omitempty"`

	// **参数解释**： 是否来自CodeHub代码仓。 **取值范围**： - true：来自CodeHub代码仓。 - false：非来自CodeHub代码仓。
	FromGitCode *bool `json:"from_git_code,omitempty"`

	// **参数解释**： 是否来自CodeHub代码仓库。 **取值范围**： - true：来自CodeHub代码仓库。 - false：非来自CodeHub代码仓库。
	FromGitCodeRepo *bool `json:"from_git_code_repo,omitempty"`

	// **参数解释**： CodeHub代码仓库ID。 **取值范围**： 不涉及。
	GitCodeRepoId *string `json:"git_code_repo_id,omitempty"`

	// **参数解释**： YAML格式流水线定义。 **取值范围**： 不涉及。
	YamlDefinition *string `json:"yaml_definition,omitempty"`

	// **参数解释**： PAC代码仓关联信息。 **取值范围**： 不涉及。
	PacRepoRelation *interface{} `json:"pac_repo_relation,omitempty"`

	// **参数解释**： YAML流水线文件内容。 **取值范围**： 不涉及。
	YamlContent *string `json:"yaml_content,omitempty"`

	// **参数解释**： 委托名称。 **取值范围**： 不涉及。
	AgencyName *string `json:"agency_name,omitempty"`

	// **参数解释**： 执行计划列表。 **取值范围**： 不涉及。
	ExecutionPlans *[]interface{} `json:"execution_plans,omitempty"`

	// **参数解释**： 流水线来源。 **取值范围**： - 0：默认。 - 1：普通模板创建。 - 2：老数据转换。 - 3：CloudInit凤凰商城触发模板创建。 - 4：CloudInit其他触发模板创建。 - 5：创建模板。
	FromSource *int32 `json:"from_source,omitempty"`

	// **参数解释**： 项目名称。 **取值范围**： 不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释**： 流水线所属分组名称。 **取值范围**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	ConcurrencyControl *PipelineConcurrencyMgmt `json:"concurrency_control,omitempty"`

	// **参数解释**： 流水线取消运行策略。 **取值范围**： 不涉及。
	CancelStrategy *interface{} `json:"cancel_strategy,omitempty"`

	// **参数解释**： 流水线标签ID列表。 **取值范围**： 不涉及。
	TagIds *[]string `json:"tag_ids,omitempty"`

	// **参数解释**： 流水线变量组列表。 **取值范围**： 不涉及。
	VariableGroups *[]string `json:"variable_groups,omitempty"`

	// **参数解释**： 流水线密级代码。 **取值范围**： 不涉及。
	SecurityLevelCode *string `json:"security_level_code,omitempty"`

	// **参数解释**： 流水线权限信息。 **取值范围**： 不涉及。
	Permissions *interface{} `json:"permissions,omitempty"`

	// **参数解释**： 主体ID，即流水线ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	SubjectId *string `json:"subject_id,omitempty"`

	// **参数解释**： 流水线详情页URL。 **取值范围**： 不涉及。
	DetailUrl *string `json:"detail_url,omitempty"`

	// **参数解释**： 流水线编辑页URL。 **取值范围**： 不涉及。
	ModifyUrl *string `json:"modify_url,omitempty"`

	// **参数解释**： 流水线标签列表。 **取值范围**： 不涉及。
	Tags *[]interface{} `json:"tags,omitempty"`

	// **参数解释**： 是否为CR（变更）模型流水线。 **取值范围**： - true：是CR模型流水线。 - false：非CR模型流水线。
	IsCrModel *bool `json:"is_cr_model,omitempty"`

	// **参数解释**： PAC归档源信息。 **取值范围**： 不涉及。
	ArchiveSource *interface{} `json:"archive_source,omitempty"`

	// **参数解释**： V2 YAML流水线的代码仓相关信息。 **取值范围**： 不涉及。
	YamlRepoProperties *interface{} `json:"yaml_repo_properties,omitempty"`

	// **参数解释**： 关联的通用参数组ID列表。 **取值范围**： 不涉及。
	VariableGroupIds *[]string `json:"variable_group_ids,omitempty"`

	// **参数解释**： PAC代码源别名。 **取值范围**： 不涉及。
	PacSourceAlias *string `json:"pac_source_alias,omitempty"`

	// **参数解释**： PAC代码源CodeHub仓库的HTTPS端点ID。 **取值范围**： 不涉及。
	PacSourceRepoHttpsEndpoint *string `json:"pac_source_repo_https_endpoint,omitempty"`
	HttpStatusCode             int     `json:"-"`
}

func (o ShowPipelineDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineDetailResponse", string(data)}, " ")
}
