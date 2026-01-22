package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsResponse Response Object
type ListConfigurationsResponse struct {

	// **参数解释：** 总记录数。 **取值范围：** 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释：** 参数模板列表。 **取值范围：** 不涉及。
	Configurations *[]ListConfigurationsResult `json:"configurations,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResponse", string(data)}, " ")
}
