package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolTagsRequest Request Object
type BatchDeletePoolTagsRequest struct {

	// **参数解释**：资源池名称。取自资源池详情的metadata字段中的name的值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	Body *DeleteTagRequest `json:"body,omitempty"`
}

func (o BatchDeletePoolTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolTagsRequest", string(data)}, " ")
}
