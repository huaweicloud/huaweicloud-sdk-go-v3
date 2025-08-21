package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Vpc struct {
	Protected *ChangedVo `json:"protected,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o Vpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vpc struct{}"
	}

	return strings.Join([]string{"Vpc", string(data)}, " ")
}
