package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdeRepositoryPair struct {

	// **参数解释**: 仓库名称。 **约束限制**: 长度1-20。 **取值范围**: 不涉及。 **默认取值**: 无。
	RepoName string `json:"repo_name"`

	// **参数解释**: 路径包含规则。 **约束限制**: 最大长度512。 **取值范围**: 不涉及。 **默认取值**: 无。
	IncludesPattern string `json:"includes_pattern"`

	// **参数解释**: 项目ID，可以从调用API处获取，也可以从控制台获取。获取方式请参考[获取项目ID](CloudArtifact_api_0015.xml)。 **约束限制**: 只能由英文字母、数字组成，且长度为32个字符。 **取值范围**: 不涉及。 **默认取值**: 无。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 仓库描述。 **约束限制**: 最大长度200。 **取值范围**: 不涉及。 **默认取值**: 无。
	Description *string `json:"description,omitempty"`

	// **参数解释**: snapshot仓库名称。 **约束限制**: 长度1-20。 **取值范围**: 不涉及。 **默认取值**: 无。
	Snapshot *string `json:"snapshot,omitempty"`

	// **参数解释**: release仓库名称。 **约束限制**: 长度1-20。 **取值范围**: 不涉及。 **默认取值**: 无。
	Release *string `json:"release,omitempty"`
}

func (o IdeRepositoryPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdeRepositoryPair struct{}"
	}

	return strings.Join([]string{"IdeRepositoryPair", string(data)}, " ")
}
