package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReviewDataInfoResponse **参数解释：** 任务门禁配置列表。
type ReviewDataInfoResponse struct {

	// **参数解释：** 运算逻辑。 **取值范围：** 取值为\"≤\"。
	CompareType *string `json:"compareType,omitempty"`

	// **参数解释：** 门禁开关。 **取值范围：** - 0：关闭门禁。 - 1：打开门禁。
	IsCheck *int32 `json:"isCheck,omitempty"`

	// **参数解释：** 门禁项。 **取值范围：** 致命问题数。 严重问题数。 一般问题数。 提示问题数。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 门禁阈值。 **取值范围：** ≥0
	Value *int32 `json:"value,omitempty"`
}

func (o ReviewDataInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewDataInfoResponse struct{}"
	}

	return strings.Join([]string{"ReviewDataInfoResponse", string(data)}, " ")
}
