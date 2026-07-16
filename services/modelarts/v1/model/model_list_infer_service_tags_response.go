package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferServiceTagsResponse Response Object
type ListInferServiceTagsResponse struct {

	// **参数解释：** 标签的融合结构，相同key合并。
	Tags           *[]CombineInferTmsTags `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListInferServiceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferServiceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInferServiceTagsResponse", string(data)}, " ")
}
