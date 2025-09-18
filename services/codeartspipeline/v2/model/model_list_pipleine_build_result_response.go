package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipleineBuildResultResponse Response Object
type ListPipleineBuildResultResponse struct {

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **取值范围**： offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每次查询的条目数量。 **取值范围**： 大于等于0。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 总条目数量。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 执行状况数据列表。 **取值范围**： 不涉及。
	BuildResults   *[]PipelineBuildResult `json:"build_results,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListPipleineBuildResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipleineBuildResultResponse struct{}"
	}

	return strings.Join([]string{"ListPipleineBuildResultResponse", string(data)}, " ")
}
