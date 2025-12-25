package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAllRepositoriesResult struct {

	// **参数解释**： 总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 仓库详情列表。 **取值范围**： 不涉及。
	Repositories *[]RepositoryBasicDo `json:"repositories,omitempty"`
}

func (o ListAllRepositoriesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllRepositoriesResult struct{}"
	}

	return strings.Join([]string{"ListAllRepositoriesResult", string(data)}, " ")
}
