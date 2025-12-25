package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectUsersRequest Request Object
type ListProjectUsersRequest struct {

	// **参数解释**: 项目ID，可以从调用API处获取，也可以从控制台获取。获取方式请参考[获取项目ID](CloudArtifact_api_0015.xml)。 **约束限制**: 只能由英文字母、数字组成，且长度为32个字符。 **取值范围**: 不涉及。 **默认取值**: 无。
	ProjectId string `json:"project_id"`

	// **参数解释**: 仓库id。可在私有库仓库**概览**界面查看。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	RepoId string `json:"repo_id"`

	// **参数解释**: scene。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	Scene *string `json:"scene,omitempty"`

	// **参数解释**: 分页查询的页数。 **约束限制**: 不涉及。 **取值范围**: 最小值1，最大值2147483647。 **默认取值**: 1
	PageNo *int32 `json:"page_no,omitempty"`

	// **参数解释**: 分页查询的每页数据量。 **约束限制**: 不涉及。 **取值范围**: 最小值1，最大值100。 **默认取值**: 10
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListProjectUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectUsersRequest struct{}"
	}

	return strings.Join([]string{"ListProjectUsersRequest", string(data)}, " ")
}
