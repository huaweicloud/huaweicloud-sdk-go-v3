package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTagResponse **参数解释**：TMS的标签结构体。
type TmsTagResponse struct {

	// **参数解释**：TMS标签的key。 **取值范围**：不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**：TMS标签的value。 **取值范围**：不涉及。
	Value *string `json:"value,omitempty"`
}

func (o TmsTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTagResponse struct{}"
	}

	return strings.Join([]string{"TmsTagResponse", string(data)}, " ")
}
