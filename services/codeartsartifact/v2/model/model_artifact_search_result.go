package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArtifactSearchResult struct {

	// **参数解释**： 文件名。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 文件相对路径。 **取值范围**： 不涉及。
	RelativePath *string `json:"relativePath,omitempty"`

	// **参数解释**： 仓库ID。 **取值范围**： 不涉及。
	Repo *string `json:"repo,omitempty"`

	// **参数解释**： 仓库名。 **取值范围**： 不涉及。
	RepoName *string `json:"repoName,omitempty"`

	// **参数解释**： 展示名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"displayName,omitempty"`

	// **参数解释**： 制品类型。 **取值范围**： 不涉及。
	RepoType *string `json:"repoType,omitempty"`

	// **参数解释**： 创建人ID。 **取值范围**： 不涉及。
	CreatedBy *string `json:"createdBy,omitempty"`

	// **参数解释**： 创建人名称。 **取值范围**： 不涉及。
	CreatedUserName *string `json:"createdUserName,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释**： 修改时间。 **取值范围**： 不涉及。
	Modified *string `json:"modified,omitempty"`

	// **参数解释**： 旧仓库ID。 **取值范围**： 不涉及。
	OldRepoId *string `json:"oldRepoId,omitempty"`
}

func (o ArtifactSearchResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArtifactSearchResult struct{}"
	}

	return strings.Join([]string{"ArtifactSearchResult", string(data)}, " ")
}
