package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Eip struct {
	Protected *ChangedVo `json:"protected,omitempty"`

	// **参数解释**： EIP数量 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o Eip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Eip struct{}"
	}

	return strings.Join([]string{"Eip", string(data)}, " ")
}
