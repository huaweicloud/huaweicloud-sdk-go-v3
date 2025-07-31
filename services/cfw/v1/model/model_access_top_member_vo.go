package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessTopMemberVo struct {

	// **参数解释**： 次数 **取值范围**： 不涉及
	Count *string `json:"count,omitempty"`

	// **参数解释**： 项 **取值范围**： 不涉及
	Item *string `json:"item,omitempty"`

	// **参数解释**： 项名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o AccessTopMemberVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessTopMemberVo struct{}"
	}

	return strings.Join([]string{"AccessTopMemberVo", string(data)}, " ")
}
