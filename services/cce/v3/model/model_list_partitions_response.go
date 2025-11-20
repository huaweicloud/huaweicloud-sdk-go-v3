package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartitionsResponse Response Object
type ListPartitionsResponse struct {

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： List
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	Items          *[]Partition `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPartitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartitionsResponse struct{}"
	}

	return strings.Join([]string{"ListPartitionsResponse", string(data)}, " ")
}
