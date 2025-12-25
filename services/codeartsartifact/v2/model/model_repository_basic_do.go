package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryBasicDo struct {

	// **参数解释**： ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 仓库状态。 **取值范围**： - active：正常。 - delete：删除。 - disabled：无效。 - view：可浏览。 - trash：废弃。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	DomainId *string `json:"domainId,omitempty"`

	// **参数解释**： 区域。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 创建时间，时间格式：yyyy-MM-dd HH:mm:ss。 **取值范围**： 不涉及。
	CreatedTime *string `json:"createdTime,omitempty"`

	// **参数解释**： 修改时间，时间格式：yyyy-MM-dd HH:mm:ss。 **取值范围**： 不涉及。
	ModifiedTime *string `json:"modifiedTime,omitempty"`

	// **参数解释**： 创建人ID。 **取值范围**： 不涉及。
	CreatedUserId *string `json:"createdUserId,omitempty"`

	// **参数解释**： 创建人名称。 **取值范围**： 不涉及。
	CreatedUserName *string `json:"createdUserName,omitempty"`

	// **参数解释**： 修改人ID。 **取值范围**： 不涉及。
	ModifiedUserId *string `json:"modifiedUserId,omitempty"`

	// **参数解释**： 修改人名称。 **取值范围**： 不涉及。
	ModifiedUserName *string `json:"modifiedUserName,omitempty"`

	// **参数解释**： 仓库ID。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 是否禁用。 **取值范围**： 不涉及。
	Disable *bool `json:"disable,omitempty"`

	// **参数解释**： 制品类型。 **取值范围**： - maven - maven2 - npm - go - pypi - rpm - composer - debian - conan - nuget - docker2 - cocoapods - ohpm - generic - helm - conda - huggingfaceml
	Format *string `json:"format,omitempty"`

	// **参数解释**： 仓库类型。 **取值范围**： - hosted：本地仓库。 - remote：代理仓库。 - virtual：虚拟仓库。
	RepoType *string `json:"repoType,omitempty"`

	// **参数解释**： 仓库策略。 **取值范围**： - release：正式仓库。 - snapshot：快照仓库。
	Policy *string `json:"policy,omitempty"`

	// **参数解释**： 用于标记一对Maven仓库(release、snapshot)，相同的tab_id即为一对Maven仓库。 **取值范围**： 不涉及。
	TabId *string `json:"tabId,omitempty"`

	// **参数解释**： 仓库名称。 **取值范围**： 不涉及。
	RepositoryName *string `json:"repositoryName,omitempty"`

	// **参数解释**： 展示的仓库名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"displayName,omitempty"`

	// **参数解释**： 仓库描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 项目id。 **取值范围**： 不涉及。
	ProjectId *string `json:"projectId,omitempty"`

	// **参数解释**： 路径包含规则。 **取值范围**： 不涉及。
	IncludesPattern *string `json:"includesPattern,omitempty"`

	// **参数解释**： 覆盖策略。 **取值范围**： 不涉及。
	DeploymentPolicy *string `json:"deploymentPolicy,omitempty"`

	// **参数解释**： 共享权限级别。 **取值范围**： PROJECT
	ShareRight *string `json:"shareRight,omitempty"`

	// **参数解释**： 仓库地址。 **取值范围**： 不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释**： 制品类型。 **取值范围**： - maven - maven2 - npm - go - pypi - rpm - composer - debian - conan - nuget - docker2 - cocoapods - ohpm - generic - helm - conda - huggingfaceml
	PackageType *string `json:"packageType,omitempty"`

	StorageSummaryInfo *DownloadFolderInfo `json:"storageSummaryInfo,omitempty"`
}

func (o RepositoryBasicDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryBasicDo struct{}"
	}

	return strings.Join([]string{"RepositoryBasicDo", string(data)}, " ")
}
