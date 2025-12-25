package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryBuildVo struct {

	// **参数解释**： 账号。 **取值范围**： 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释**： 密码。 **取值范围**： 不涉及。
	Password *string `json:"password,omitempty"`

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
	Type *string `json:"type,omitempty"`

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

	// **参数解释**： snapshot仓库名称,release和snapshot至少二选一。 **取值范围**： 不涉及。
	Snapshot *string `json:"snapshot,omitempty"`

	// **参数解释**： release仓库名称,release和snapshot至少二选一。 **取值范围**： 不涉及。
	Release *string `json:"release,omitempty"`

	// **参数解释**： npm。 **取值范围**： 不涉及。
	Npm *string `json:"npm,omitempty"`

	// **参数解释**： 快照仓库状态。 **取值范围**： 不涉及。
	SnapshotStatus *string `json:"snapshotStatus,omitempty"`

	// **参数解释**： 正式仓库状态。 **取值范围**： 不涉及。
	ReleaseStatus *string `json:"releaseStatus,omitempty"`

	// **参数解释**： 项目id。 **取值范围**： 不涉及。
	ProjectId *string `json:"projectId,omitempty"`

	// **参数解释**： 路径包含规则。 **取值范围**： 不涉及。
	IncludesPattern *string `json:"includesPattern,omitempty"`

	// **参数解释**： 仓库ID列表。 **取值范围**： 不涉及。
	RepositoryIds *[]string `json:"repositoryIds,omitempty"`

	// **参数解释**： uri。 **取值范围**： 不涉及。
	Uri *string `json:"uri,omitempty"`

	// **参数解释**： 覆盖策略。 **取值范围**： 不涉及。
	DeploymentPolicy *string `json:"deploymentPolicy,omitempty"`

	// **参数解释**： 仓库列表。 **取值范围**： 不涉及。
	Repositories *[]string `json:"repositories,omitempty"`

	// **参数解释**： 父仓库名。 **取值范围**： 不涉及。
	ParentRepoName *string `json:"parentRepoName,omitempty"`

	// **参数解释**: 用户名。 **取值范围**: 不涉及。
	UserName *string `json:"userName,omitempty"`

	// **参数解释**： 代理仓地址。 **取值范围**： 不涉及。
	RemoteUrl *string `json:"remoteUrl,omitempty"`

	// **参数解释**： 默认仓库。 **取值范围**： 不涉及。
	DefaultDeployRepository *string `json:"defaultDeployRepository,omitempty"`

	// **参数解释**： 代理仓类型。 **取值范围**： - public：公共代理仓； - customize：用户自定义代理仓。
	RemoteType *string `json:"remoteType,omitempty"`

	// **参数解释**： 代理。 **取值范围**： 不涉及。
	Proxy *string `json:"proxy,omitempty"`

	// **参数解释**： 是否允许匿名下载。 **取值范围**： 不涉及。
	AllowAnonymous *bool `json:"allowAnonymous,omitempty"`

	// **参数解释**： 是否自动清理快照版本。 **取值范围**： 不涉及。
	AutoCleanSnapshot *bool `json:"autoCleanSnapshot,omitempty"`

	// **参数解释**： 快照版本有效期，单位：天。 **取值范围**： 不涉及。
	SnapshotAliveDays *string `json:"snapshotAliveDays,omitempty"`

	// **参数解释**： 最大快照数。 **取值范围**： 不涉及。
	MaxUniqueSnapshots *string `json:"maxUniqueSnapshots,omitempty"`

	// **参数解释**： 共享权限级别。 **取值范围**： PROJECT
	ShareRight *string `json:"shareRight,omitempty"`

	// **参数解释**： 是否nexus仓库。 **取值范围**： 不涉及。
	NexuRepo *bool `json:"nexuRepo,omitempty"`

	// **参数解释**： 仓库地址。 **取值范围**： 不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释**： 制品类型。 **取值范围**： - maven - maven2 - npm - go - pypi - rpm - composer - debian - conan - nuget - docker2 - cocoapods - ohpm - generic - helm - conda - huggingfaceml
	PackageType *string `json:"packageType,omitempty"`
}

func (o RepositoryBuildVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryBuildVo struct{}"
	}

	return strings.Join([]string{"RepositoryBuildVo", string(data)}, " ")
}
