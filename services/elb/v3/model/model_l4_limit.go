package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type L4Limit struct {

	// **参数解释**：负载均衡实例的四层规格的最大并发连接数限速值。  **取值范围**：不涉及
	Connection *int32 `json:"connection,omitempty"`

	// **参数解释**：负载均衡实例的四层规格的每秒新建连接数限速值。  **取值范围**：不涉及
	Cps *int32 `json:"cps,omitempty"`
}

func (o L4Limit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L4Limit struct{}"
	}

	return strings.Join([]string{"L4Limit", string(data)}, " ")
}
