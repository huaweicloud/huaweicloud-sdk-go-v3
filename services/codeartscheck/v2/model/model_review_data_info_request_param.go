package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReviewDataInfoRequestParam **参数解释：** 任务门禁配置列表。 **约束限制：** 不涉及。
type ReviewDataInfoRequestParam struct {

	// **参数解释：** 运算逻辑。 **约束限制：** 不涉及。 **取值范围：** 取值为\"≤\"。 **默认取值：** 不涉及。
	CompareType string `json:"compareType"`

	// **参数解释：** 门禁开关。 **约束限制：** 不涉及。 **取值范围：** - 0：关闭门禁。 - 1：打开门禁。 **默认取值：** 不涉及。
	IsCheck int32 `json:"isCheck"`

	// **参数解释：** 门禁项。 **约束限制：** 不涉及。 **取值范围：** - 致命问题数。 - 严重问题数。 - 一般问题数。 - 提示问题数。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 门禁阈值。 **约束限制：** 不涉及。 **取值范围：** ≥0 **默认取值：** 不涉及。
	Value int32 `json:"value"`
}

func (o ReviewDataInfoRequestParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewDataInfoRequestParam struct{}"
	}

	return strings.Join([]string{"ReviewDataInfoRequestParam", string(data)}, " ")
}
