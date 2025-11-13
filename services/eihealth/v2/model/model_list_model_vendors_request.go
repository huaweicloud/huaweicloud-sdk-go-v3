package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelVendorsRequest Request Object
type ListModelVendorsRequest struct {

	// **参数解释**： 供应商类型。 **约束限制**： 不涉及 **取值范围**： * system：系统内置 * custom：自定义 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 排序规则，目前默认创建时间降序。 **约束限制**： 不涉及 **取值范围**： - name：根据创建时间排列。 - status：模型服务状态 - update_time：根据更新时间排列。 **默认取值**： update_time
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**： 排序方向。 **约束限制**： 不涉及 **取值范围**： - acs：升序 - desc：降序 **默认取值**： desc
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**： 限制量，单次查询总量。 **约束限制**： 不涉及 **取值范围**： [1-1000] **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，查询起始偏移。 **约束限制**： 不涉及 **取值范围**： [0-100000000] **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListModelVendorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelVendorsRequest struct{}"
	}

	return strings.Join([]string{"ListModelVendorsRequest", string(data)}, " ")
}
