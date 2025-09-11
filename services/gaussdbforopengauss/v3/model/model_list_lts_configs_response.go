package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsConfigsResponse Response Object
type ListLtsConfigsResponse struct {

	// **参数解释**: 总记录数。 **取值范围**: 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**: 实例下Lts明细列表。
	InstanceLtsConfigs *[]InstanceLtsLogConfigResult `json:"instance_lts_configs,omitempty"`
	HttpStatusCode     int                           `json:"-"`
}

func (o ListLtsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsConfigsResponse", string(data)}, " ")
}
