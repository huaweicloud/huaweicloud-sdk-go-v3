package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTmsTagsRequest 给资源添加标签请求体。
type CreateTmsTagsRequest struct {

	// **参数解释**：TMS标签列表。 **约束限制**：不涉及。
	Tags []TmsTag `json:"tags"`
}

func (o CreateTmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTmsTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateTmsTagsRequest", string(data)}, " ")
}
