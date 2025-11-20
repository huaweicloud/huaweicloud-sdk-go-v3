package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfo struct {

	// - 参数解释：当前页的最后一条记录，最后一页时无next_marker字段。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	NextMarker *string `json:"next_marker,omitempty"`

	// - 参数解释：分页查询起始的资源ID，为空时查询第一页。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	PreviousMarker string `json:"previous_marker"`

	// - 参数解释：当前页记录总数。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	CurrentCount int32 `json:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
