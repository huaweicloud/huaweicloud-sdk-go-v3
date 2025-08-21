package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ItemVo struct {

	// **参数解释**： 聚合项 **取值范围**： 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释**： 聚合项名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 统计值 **取值范围**： 不涉及
	Value *int64 `json:"value,omitempty"`
}

func (o ItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemVo struct{}"
	}

	return strings.Join([]string{"ItemVo", string(data)}, " ")
}
