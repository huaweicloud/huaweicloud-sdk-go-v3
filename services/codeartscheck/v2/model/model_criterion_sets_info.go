package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CriterionSetsInfo struct {

	// **参数解释：** 规则列表。 **取值范围：** 不涉及。
	CriterionSetList *[]CriterionSetListV2Body `json:"criterionSetList,omitempty"`

	// **参数解释：** 规则总数。 **取值范围：** ≥0
	Total *int32 `json:"total,omitempty"`
}

func (o CriterionSetsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CriterionSetsInfo struct{}"
	}

	return strings.Join([]string{"CriterionSetsInfo", string(data)}, " ")
}
