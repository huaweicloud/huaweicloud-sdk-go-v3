package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAttentionResult struct {

	// **参数解释**： 序号。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	DomainId *string `json:"domainId,omitempty"`

	// **参数解释**： 仓库ID。 **取值范围**： 不涉及。
	RepositoryId *string `json:"repositoryId,omitempty"`

	// **参数解释**： 仓库名称。 **取值范围**： 不涉及。
	RepositoryName *string `json:"repositoryName,omitempty"`

	// **参数解释**： 制品类型。 **取值范围**： maven2|docker|npm|go|pypi|rpm|composer|debian|conan|nuget|docker2|cocoapods|ohpm|generic。
	Format *string `json:"format,omitempty"`

	// **参数解释**： 仓库策略。 **取值范围**： - release：正式仓库。 - snapshot：快照仓库。
	Policy *string `json:"policy,omitempty"`

	// **参数解释**： 关注的组件序号。 **取值范围**： 不涉及。
	ArtifactId *string `json:"artifactId,omitempty"`

	// **参数解释**： 关注的组件路径。 **取值范围**： 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释**： 修改人名称。 **取值范围**： 不涉及。
	ModifiedUserName *string `json:"modifiedUserName,omitempty"`

	// **参数解释**： 修改人ID。 **取值范围**： 不涉及。
	ModifiedUserId *string `json:"modifiedUserId,omitempty"`

	// **参数解释**： 用户id。 **取值范围**： 只能由英文字母、数字组成，且长度为32个字符。
	UserId *string `json:"userId,omitempty"`

	// **参数解释**： 修改时间，时间格式：yyyy-MM-dd HH:mm:ss。 **取值范围**： 不涉及。
	ModifiedTime *string `json:"modifiedTime,omitempty"`

	// **参数解释**： 区域。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`
}

func (o ListAttentionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttentionResult struct{}"
	}

	return strings.Join([]string{"ListAttentionResult", string(data)}, " ")
}
