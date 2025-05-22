package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimsInfo CES监控维护信息
type DimsInfo struct {

	// **参数解释**： CES监控维度路由。 **取值范围**： 不涉及。
	DimK *string `json:"dim_k,omitempty"`

	// **参数解释**： CES监控维度ID。 **取值范围**： 不涉及。
	DimV *string `json:"dim_v,omitempty"`
}

func (o DimsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimsInfo struct{}"
	}

	return strings.Join([]string{"DimsInfo", string(data)}, " ")
}
