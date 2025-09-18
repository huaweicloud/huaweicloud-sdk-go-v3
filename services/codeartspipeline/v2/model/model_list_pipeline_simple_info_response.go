package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineSimpleInfoResponse Response Object
type ListPipelineSimpleInfoResponse struct {

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **取值范围**： 大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每次查询的条目数量。 **取值范围**： 大于等于0。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 总条目数量。 **取值范围**： 大于等于0。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 流水线列表。 **取值范围**： 不涉及。
	Result         *[]PipelineBasicInfo `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListPipelineSimpleInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineSimpleInfoResponse struct{}"
	}

	return strings.Join([]string{"ListPipelineSimpleInfoResponse", string(data)}, " ")
}
