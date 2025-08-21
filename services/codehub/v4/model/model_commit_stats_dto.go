package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitStatsDto struct {

	// **参数解释：** 增加行数。 **取值范围：** 不涉及。
	Additions *int32 `json:"additions,omitempty"`

	// **参数解释：** 删除行数。 **取值范围：** 不涉及。
	Deletions *int32 `json:"deletions,omitempty"`

	// **参数解释：** 变更行数。 **取值范围：** 不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o CommitStatsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitStatsDto struct{}"
	}

	return strings.Join([]string{"CommitStatsDto", string(data)}, " ")
}
