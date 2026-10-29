package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigTemplateBasicInfo **参数解释：** 配置模板详情。
type ConfigTemplateBasicInfo struct {

	// **参数解释：** 执行机类型。 **取值范围：** X86/ARM。
	Arch *string `json:"arch,omitempty"`

	// **参数解释：** 分支名。 **取值范围：** 不涉及。
	Branch *string `json:"branch,omitempty"`

	// **参数解释：** 配置模板创建时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：** 配置模板创建者。 **取值范围：** 不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：** 任务配置的排除目录信息。 **取值范围:** 不涉及。
	Exclude *string `json:"exclude,omitempty"`

	// **参数解释：** 配置模板ID，修改配置信息需要用到配置模板ID，每个代码检查任务及其下的分支任务都有相同的且唯一的配置模板ID。 **取值范围:** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 检查目录。不配置检查目录时，检查代码仓中所有目录下的代码文件；配置检查目录时，只检查配置目录中的文件。 **取值范围：** 不涉及。
	IncludePaths *string `json:"includePaths,omitempty"`

	// **参数解释：** 最后修改人的用户名。 **取值范围：** 不涉及。
	LastReviser *string `json:"lastReviser,omitempty"`

	// **参数解释：** 配置模板上一次修改时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：** 配置模板的名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 项目的32位UUID，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/zh-cn/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)接口获取。 **取值范围：** 长度为[1~32]个字符。
	ProjectId *string `json:"projectId,omitempty"`

	// **参数解释：** 项目名。 **取值范围：** 不涉及。
	ProjectName *string `json:"projectName,omitempty"`

	// **参数解释：** 华为云地域。 **取值范围：** 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释：** 代码仓类型。 **取值范围：** - DevCloud：CodeArtsRepo仓。 - Gitee：Gitee仓。 - GitHub：GitHub仓。 - GitPub：通用Git仓。 - gitcode：GitCode仓。 - gitlab：GitLab仓。 - self_gitlab：自建GitLab仓。 - tfs_git：TFS仓。 - bitbucket：Bitbucket仓。
	RepoType *string `json:"repoType,omitempty"`

	// **参数解释：** 代码仓地址。 **取值范围：** 不涉及。
	RepoUrl *string `json:"repoUrl,omitempty"`

	// **参数解释：** 启用的资源池ID，使用内置执行机的情况下为\"\"。 **取值范围：** 不涉及。
	ResourcePoolId *string `json:"resourcePoolId,omitempty"`

	// **参数解释：** 启用的资源池名称，使用内置执行机的情况下为空。 **取值范围：** 不涉及。
	ResourcePoolName *string `json:"resourcePoolName,omitempty"`

	// **参数解释：** 启用的资源池类型。 **取值范围：** - default：使用内置资源池。 - exclusive：使用专属资源池。 - self-hosted：使用自定义资源池。
	ResourcePoolType *string `json:"resource_pool_type,omitempty"`

	// **参数解释：** 任务门禁配置列表。
	ReviewData *[]ReviewDataInfoResponse `json:"reviewData,omitempty"`

	// **参数解释：** 任务门禁配置列表的json字符串。 **取值范围：** 不涉及。
	ReviewDataJson *string `json:"reviewDataJson,omitempty"`

	// **参数解释：** 规则集配置列表。
	RuleSets *[]ConfigTemplateRuleSetInfo `json:"ruleSets,omitempty"`
}

func (o ConfigTemplateBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigTemplateBasicInfo struct{}"
	}

	return strings.Join([]string{"ConfigTemplateBasicInfo", string(data)}, " ")
}
