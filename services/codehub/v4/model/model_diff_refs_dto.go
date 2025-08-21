package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffRefsDto struct {

	// **参数解释：** 目标分支的基准提交哈希。 **取值范围：** 不涉及。
	BaseSha *string `json:"base_sha,omitempty"`

	// **参数解释：** 源分支的最新提交哈希。 **取值范围：** 不涉及。
	HeadSha *string `json:"head_sha,omitempty"`

	// **参数解释：** 合并请求开始时的提交哈希，通常与base_sha相同。 **取值范围：** 不涉及。
	StartSha *string `json:"start_sha,omitempty"`
}

func (o DiffRefsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffRefsDto struct{}"
	}

	return strings.Join([]string{"DiffRefsDto", string(data)}, " ")
}
