package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNotMavenRepoDo struct {

	// **参数解释**: 仓库名称。 **约束限制**: 长度1-20。 **取值范围**: 不涉及。 **默认取值**: 无。
	RepoName string `json:"repo_name"`

	// **参数解释**: 制品类型。 **约束限制**: 不涉及。 **取值范围**: docker|npm|go|pypi|rpm|composer|debian|conan|nuget|docker2|cocoapods|ohpm|generic。 **默认取值**: 无。
	Format string `json:"format"`

	// **参数解释**: 仓库描述。 **约束限制**: 最大长度200。 **取值范围**: 不涉及。 **默认取值**: 无。
	Description *string `json:"description,omitempty"`

	// **参数解释**: 仓库id列表。仓库id，格式为{region}_{domainId}_{format}_{sequence}。可以从私有依赖库首页->仓库概览->仓库地址 url 中获取，最后两个\"/\"中间的字符串即为仓库id。 **约束限制**: 根据仓库id格式中region, domainId需要为有效值，format有效值为:npm|go|pypi|rpm|composer|maven|debian|conan|nuget|docker2|cocoapods|ohpm, sequence取值根据套餐不同有不同上限值。 **取值范围**: 不涉及。 **默认取值**: 无。
	RepositoryIds []string `json:"repository_ids"`

	// **参数解释**: 路径包含规则。 **约束限制**: 最大长度512。 **取值范围**: 不涉及。 **默认取值**: 无。
	IncludesPattern string `json:"includes_pattern"`

	// **参数解释**: 覆盖策略。 **约束限制**: 不涉及。 **取值范围**: allowRedeploy：允许覆盖 disableRedeploy：禁止覆盖 readOnly：只读。 **默认取值**: 无。
	DeploymentPolicy *string `json:"deployment_policy,omitempty"`

	// **参数解释**: 自动清理快照。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	AutoCleanSnapshot *bool `json:"auto_clean_snapshot,omitempty"`

	// **参数解释**: 快照保存时间长度。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	SnapshotAliveDays *string `json:"snapshot_alive_days,omitempty"`

	// **参数解释**: 最大不同快照个数。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	MaxUniqueSnapshots *string `json:"max_unique_snapshots,omitempty"`

	// **参数解释**: 是否允许匿名。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	AllowAnonymous *bool `json:"allow_anonymous,omitempty"`

	// **参数解释**: 项目ID，可以从调用API处获取，也可以从控制台获取。获取方式请参考[获取项目ID](CloudArtifact_api_0015.xml)。 **约束限制**: 只能由英文字母、数字组成，且长度为32个字符。 **取值范围**: 不涉及。 **默认取值**: 无。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o UpdateNotMavenRepoDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotMavenRepoDo struct{}"
	}

	return strings.Join([]string{"UpdateNotMavenRepoDo", string(data)}, " ")
}
