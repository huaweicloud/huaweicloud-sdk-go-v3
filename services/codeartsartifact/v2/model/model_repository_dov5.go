package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryDov5 struct {

	// **参数解释**: 仓库状态。 **取值范围**: active：正常。 delete：删除。 disabled：无效。 view：私有库浏览者。 trash：废弃。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 租户id。 **取值范围**: 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**: 区域。 **取值范围**: 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**: 创建时间，时间格式：yyyy-MM-dd HH:mm:ss。 **取值范围**: 不涉及。
	CreatedTime *string `json:"created_time,omitempty"`

	// **参数解释**: 修改时间，时间格式：yyyy-MM-dd HH:mm:ss。 **取值范围**: 不涉及。
	ModifiedTime *string `json:"modified_time,omitempty"`

	// **参数解释**: 创建人id。 **取值范围**: 不涉及。
	CreatedUserId *string `json:"created_user_id,omitempty"`

	// **参数解释**: 创建人。 **取值范围**: 不涉及。
	CreatedUserName *string `json:"created_user_name,omitempty"`

	// **参数解释**: 修改人id。 **取值范围**: 不涉及。
	ModifiedUserId *string `json:"modified_user_id,omitempty"`

	// **参数解释**: 修改人。 **取值范围**: 不涉及。
	ModifiedUserName *string `json:"modified_user_name,omitempty"`

	// **参数解释**: 仓库名称。 **取值范围**: 不涉及。
	RepositoryName *string `json:"repository_name,omitempty"`

	// **参数解释**: 制品类型。 **取值范围**: - maven - maven2 - npm - go - pypi - rpm - composer - debian - conan - nuget - docker2 - cocoapods - ohpm - generic - helm - conda
	Format *string `json:"format,omitempty"`

	// **参数解释**: 仓库类型。 **取值范围**: hosted：本地仓库。 remote：代理仓库。 virtual：聚合仓库。
	Type *string `json:"type,omitempty"`

	// **参数解释**: 仓库描述。 **取值范围**: 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**: release仓库名称,release和snapshot至少二选一。 **取值范围**: 不涉及。
	Release *string `json:"release,omitempty"`

	// **参数解释**: snapshot仓库名称,release和snapshot至少二选一。 **取值范围**: 不涉及。
	Snapshot *string `json:"snapshot,omitempty"`

	// **参数解释**: 路径包含规则。 **取值范围**: 不涉及。
	IncludesPattern *string `json:"includes_pattern,omitempty"`

	// **参数解释**: 路径排除规则。 **取值范围**: 不涉及。
	ExcludesPattern *string `json:"excludes_pattern,omitempty"`

	// **参数解释**: 共享权限级别。 **取值范围**: PROJECT。
	ShareRight *string `json:"share_right,omitempty"`

	// **参数解释**: 项目ID。 **取值范围**: 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 仓库id。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 是否禁用。 **取值范围**: 不涉及。
	Disable *bool `json:"disable,omitempty"`

	// **参数解释**: 仓库策略。 **取值范围**: release或snapshot。
	Policy *string `json:"policy,omitempty"`

	// **参数解释**: npm。 **取值范围**: 不涉及。
	Npm *string `json:"npm,omitempty"`

	// **参数解释**: uri。 **取值范围**: 不涉及。
	Uri *string `json:"uri,omitempty"`

	// **参数解释**: repositories。 **取值范围**: 不涉及。
	Repositories *string `json:"repositories,omitempty"`

	// **参数解释**: 账号。 **取值范围**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 密码。 **取值范围**: 不涉及。
	Password *string `json:"password,omitempty"`

	// **参数解释**: 代理。 **取值范围**: 不涉及。
	Proxy *string `json:"proxy,omitempty"`

	// **参数解释**: 范围。 **取值范围**: 不涉及。
	Scope *int32 `json:"scope,omitempty"`

	// **参数解释**: 地址。 **取值范围**: 不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释**: 用于标记一对maven仓库(release、snapshot)，相同的tab_id即为一对maven仓库。 **取值范围**: 不涉及。
	TabId *string `json:"tab_id,omitempty"`

	// **参数解释**: 展示的仓库名称。 **取值范围**: 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**: 快照仓状态。 **取值范围**: 不涉及。
	SnapshotStatus *string `json:"snapshot_status,omitempty"`

	// **参数解释**: 发布仓状态。 **取值范围**: 不涉及。
	ReleaseStatus *string `json:"release_status,omitempty"`

	// **参数解释**: 仓库id列表。 **取值范围**: 不涉及。
	RepositoryIds *string `json:"repository_ids,omitempty"`

	// **参数解释**: 覆盖策略。 **取值范围**: 不涉及。
	DeploymentPolicy *string `json:"deployment_policy,omitempty"`

	// **参数解释**: 父仓库名。 **取值范围**: 不涉及。
	ParentRepoName *string `json:"parent_repo_name,omitempty"`

	// **参数解释**: 代理仓地址。 **取值范围**: 不涉及。
	RemoteUrl *string `json:"remote_url,omitempty"`

	// **参数解释**: 代理仓认证。 **取值范围**: 不涉及。
	RemoteAuth *string `json:"remote_auth,omitempty"`

	// **参数解释**: pypi索引代理地址。 **取值范围**: 不涉及。
	PypiRegistryUrl *string `json:"pypi_registry_url,omitempty"`

	// **参数解释**: 虚仓的默认仓库。 **取值范围**: 不涉及。
	DefaultDeployRepository *string `json:"default_deploy_repository,omitempty"`

	// **参数解释**: 制品类型。 **取值范围**: 不涉及。
	PackageType *string `json:"package_type,omitempty"`

	// **参数解释**: 是否nexu仓库。 **取值范围**: 不涉及。
	NexuRepo *bool `json:"nexu_repo,omitempty"`

	// **参数解释**: 迁移标识。 **取值范围**: 不涉及。
	MigrateFlag *int32 `json:"migrate_flag,omitempty"`
}

func (o RepositoryDov5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryDov5 struct{}"
	}

	return strings.Join([]string{"RepositoryDov5", string(data)}, " ")
}
