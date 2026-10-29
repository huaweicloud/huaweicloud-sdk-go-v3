package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskBasicInfo **参数解释：** 任务信息。 **取值范围：** 不涉及。
type TaskBasicInfo struct {

	// **参数解释：** 分支名称。 **取值范围：** 不涉及。
	BranchName *string `json:"branchName,omitempty"`

	// **参数解释：** 是否有删除任务的权限。 **取值范围：** - true：有删除的权限。 - false：没有删除的权限。
	CanDelete *bool `json:"canDelete,omitempty"`

	// **参数解释：** 是否有编辑的权限。 **取值范围：** - true：有编辑的权限。 - false：没有编辑的权限。
	CanEdit *bool `json:"canEdit,omitempty"`

	// **参数解释：** 是否有执行的权限。 **取值范围：** - true：有执行的权限。 - false：没有执行的权限。
	CanExecute *bool `json:"canExecute,omitempty"`

	// **参数解释：** 是否有读的权限。 **取值范围：** - true：有读的权限。 - false：没有读的权限。
	CanRead *bool `json:"canRead,omitempty"`

	ConfigTemplate *ConfigTemplateBasicInfo `json:"configTemplate,omitempty"`

	// **参数解释：** 任务的配置模板的ID。 **取值范围：** 长度为[1~128]个字符。
	ConfigTemplateId *string `json:"configTemplateId,omitempty"`

	// **参数解释：** 任务的创建时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：** 任务ID，即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。即[[查询任务列表](https://support.huaweicloud.com/api-codecheck/ShowTasksList.html)](tag:hws)[[查询任务列表](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ShowTasksList.html)](tag:hws_hk)[查询任务列表](tag:hcs,hcs_sm)接口响应中的id字段。 **取值范围：** 长度为[1~128]个字符。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 任务是否被关注。 **取值范围：** - true：任务被关注。 - false：任务未被关注。
	IsFollowed *bool `json:"isFollowed,omitempty"`

	JobResult *JobResultInfo `json:"jobResult,omitempty"`

	LastJob *JobBasicInfo `json:"lastJob,omitempty"`

	// **参数解释：** 任务的上一次更新时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：** 代码检查任务名。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 代码检查任务所有者。 **取值范围：** 不涉及。
	Owner *string `json:"owner,omitempty"`

	// **参数解释：** 项目的32位UUID，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/zh-cn/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)接口获取。 **取值范围：** 长度为[1~32]个字符。
	ProjectId *string `json:"projectId,omitempty"`

	// **参数解释：** 项目名。 **取值范围：** 不涉及。
	ProjectName *string `json:"projectName,omitempty"`

	// **参数解释：** 华为云地域。 **取值范围：** 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释：** 代码仓地址。 **取值范围：** 不涉及。
	RepoUrl *string `json:"repoUrl,omitempty"`

	// **参数解释：** 规则集配置列表。
	RuleSets *[]ConfigTemplateRuleSetInfo `json:"ruleSets,omitempty"`

	// **参数解释：** 任务类型。 **取值范围：** - basic：普通任务。 - security：安全增强包任务。
	RuleType *string `json:"ruleType,omitempty"`

	// **参数解释：** 任务分组的ID。 **取值范围：** 长度为[1~128]个字符。
	TaskGroupId *string `json:"taskGroupId,omitempty"`
}

func (o TaskBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBasicInfo struct{}"
	}

	return strings.Join([]string{"TaskBasicInfo", string(data)}, " ")
}
