package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metadata **参数解释：** 返回满足过滤条件的资源总数。 **取值范围：** 不涉及。
type Metadata struct {

	// **参数解释：** 满足查询条件的资源总数，不受分页（即limit、offset参数）影响。 **取值范围：** 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
