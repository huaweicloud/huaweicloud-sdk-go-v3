package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodelineDto 仓库近15日每日代码提交行数增减信息
type CodelineDto struct {

	// **参数解释：** 增加行数。 **取值范围：** 最小0 **默认取值：** 0
	Additions *int32 `json:"additions,omitempty"`

	// **参数解释：** 删除行数。 **取值范围：** 最小0 **默认取值：** 0
	Deletions *int32 `json:"deletions,omitempty"`

	// **参数解释：** 日期，格式'yyyyMMdd',举例：20251030。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Date *string `json:"date,omitempty"`
}

func (o CodelineDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodelineDto struct{}"
	}

	return strings.Join([]string{"CodelineDto", string(data)}, " ")
}
