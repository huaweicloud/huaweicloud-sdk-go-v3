package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTagResp TMS的标签结构体。
type TmsTagResp struct {

	// **参数解释**：TMS标签的key。 **取值范围**：不涉及。
	Key string `json:"key"`

	// **参数解释**：TMS标签的value。 **取值范围**：不涉及。
	Value string `json:"value"`
}

func (o TmsTagResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTagResp struct{}"
	}

	return strings.Join([]string{"TmsTagResp", string(data)}, " ")
}
