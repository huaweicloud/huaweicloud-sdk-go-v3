package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopInfo struct {

	// **参数解释**： 次数 **取值范围**： 不涉及
	Count *int64 `json:"count,omitempty"`

	// **参数解释**： 项 **取值范围**： 不涉及
	Item *string `json:"item,omitempty"`

	// **参数解释**： 项ID **取值范围**： 不涉及
	ItemId *string `json:"item_id,omitempty"`
}

func (o TopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopInfo struct{}"
	}

	return strings.Join([]string{"TopInfo", string(data)}, " ")
}
